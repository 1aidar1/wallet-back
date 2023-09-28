package entity

// Currency
type Currency struct {
	ID   string
	Code string
}

var EmptyCurrency = Currency{}

// Country
type Country struct {
	ID   string
	Code string
}

var EmptyCountry = Country{}

// WalletStatus
type WalletStatus string

func (s WalletStatus) String() string {
	return string(s)
}

// OperationStatus
type OperationStatus string

func (s OperationStatus) String() string {
	return string(s)
}

// OperationType
type OperationType string

func (s OperationType) String() string {
	return string(s)
}

// TransactionType
type TransactionType string

func (s TransactionType) String() string {
	return string(s)
}

// IdentificationType none, full, basic
type IdentificationType string

func (s IdentificationType) String() string {
	return string(s)
}
