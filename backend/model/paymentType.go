package model

//Enum for allowed payment types
type PaymentType int

const (
	CreditCard PaymentType = iota
)

func (pt PaymentType) String() string {
	switch pt {
	case CreditCard:
		return "Credit Card"
	}
	return "Unknown"
}
