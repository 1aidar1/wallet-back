package dictionary

import "git.example.kz/wallet/wallet-back/internal/entity"

type WalletStatuses struct {
	Active  entity.WalletStatus
	Blocked entity.WalletStatus
	Closed  entity.WalletStatus
}

func initWalletStatusesEnums() WalletStatuses {
	//enum'ы с базы
	statuses := WalletStatuses{
		Active:  "active",
		Blocked: "blocked",
		Closed:  "closed",
	}

	return statuses
}
