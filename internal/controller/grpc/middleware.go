package grpc_server

import (
	"context"
	"fmt"
	"runtime/debug"

	"git.example.kz/wallet/wallet-back/pkg/errcode"
	"git.example.kz/wallet/wallet-back/pkg/logger"
	"google.golang.org/grpc"
)

// Recover returns a new unary server interceptor for panic recovery.
func Recover() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = errcode.NewDefaultErr()
				logger.GetProdLogger().Error(fmt.Sprintf("Panicked recovered in grpc interceptor. Msg: %s Stack: %s", r, debug.Stack()))
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false
		return resp, err
	}
}
