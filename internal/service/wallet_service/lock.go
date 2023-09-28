package wallet_service

import (
	"context"
	"fmt"
	"time"

	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"github.com/redis/go-redis/v9"
)

// Подождать пока спадет лок, после залочить на новую операцию.
// если айдишников несколько, ждем пока все они будут свободны.
func (s *WalletService) WaitAndLock(ctx context.Context, walletIds ...string) *errcode.ErrCode {

	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop() // Stop the ticker to release resources when done
	for {
		select {
		case <-ctx.Done():
			return errcode.New("wallet_locked")
		case <-ticker.C:
			unlocked, err := s.AreUnlocked(ctx, walletIds...)
			if err != nil {
				return err
			}
			if unlocked {
				locked, err := s.Lock(ctx, walletIds...)
				if err != nil {
					return err
				}
				if locked {
					return nil
				}
			}
		}
	}

}

func (s *WalletService) Lock(ctx context.Context, walletIds ...string) (bool, *errcode.ErrCode) {
	// Start a Redis transaction

	for _, walletId := range walletIds {
		b, err := s.redis.SetNX(ctx, lockKey(walletId), true, s.lockTime).Result()
		if err != nil {
			e := errcode.New(errcode.DefaultCode).WithErr(err)
			logger.GetProdLogger().Error(fmt.Sprintf("couldnt lock wallet in redis %s | %s", walletIds, e.Log()))
			return false, e
		}
		if !b {
			return false, nil
		}
	}

	return true, nil
}

func (s *WalletService) AreUnlocked(ctx context.Context, walletIds ...string) (bool, *errcode.ErrCode) {
	// Start a Redis transaction
	tx := s.redis.TxPipeline()

	// Add EXISTS commands for each wallet WalletId to the transaction
	existsCmds := make([]*redis.IntCmd, len(walletIds))
	for i, walletId := range walletIds {
		existsCmds[i] = tx.Exists(ctx, lockKey(walletId))
	}

	// Execute the transaction and get the results of the EXISTS commands
	_, err := tx.Exec(ctx)
	if err != nil {
		e := errcode.New(errcode.DefaultCode).WithErr(err)
		logger.GetProdLogger().Error(fmt.Sprintf("couldnt scan if wallet is blocked from redis %s | %s", walletIds, e.Log()))
		return false, e
	}

	// Check if all keys exist
	for _, cmd := range existsCmds {
		if cmd.Val() == 1 {
			return false, nil
		}
	}

	return true, nil
}

func (s *WalletService) Unlock(ctx context.Context, walletIds ...string) *errcode.ErrCode {
	// Start a Redis transaction
	tx := s.redis.TxPipeline()

	// Add DEL commands to unlock each wallet WalletId
	for _, walletId := range walletIds {
		tx.Del(ctx, lockKey(walletId))
	}

	// Execute the transaction
	_, err := tx.Exec(ctx)
	if err != nil {
		e := errcode.New(errcode.DefaultCode).WithErr(err)
		logger.GetProdLogger().Error(fmt.Sprintf("couldnt unlock wallet in redis %s | %s", walletIds, e.Log()))
		return e
	}

	return nil
}

func lockKey(id string) string {
	return fmt.Sprintf("wallet_lock:%s", id)
}
