package model

//Enum for allowed payment types
type PaymentType int

const (
	Blik PaymentType = iota
	CreditCard
	BankTransfer
)

func (pt PaymentType) String() string {
	switch pt {
	case Blik:
		return "BLIK"
	case CreditCard:
		return "Credit Card"
	case BankTransfer:
		return "Bank Transfer"
	}
	return "Unknown"
}