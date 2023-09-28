package dictionary

import "git.example.kz/wallet/wallet-back/internal/entity"

type WalletIdentification struct {
	Full  entity.IdentificationType
	Basic entity.IdentificationType
	None  entity.IdentificationType
}

func initWalletIdentificationEnums() WalletIdentification {
	//enum'ы с базы
	return WalletIdentification{
		Full:  "full",
		Basic: "basic",
		None:  "none",
	}

}
