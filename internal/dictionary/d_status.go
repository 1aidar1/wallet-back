package dictionary

import "git.example.kz/wallet/wallet-back/internal/entity"

type Statuses struct {
	Process entity.OperationStatus
	Success entity.OperationStatus
	Error   entity.OperationStatus
}

func initStatusesEnums() Statuses {
	statuses := Statuses{
		Process: "process",
		Success: "success",
		Error:   "error",
	}

	return statuses
}
