package rest

import (
	"github.com/RobertOchmanek/ebiznes_go/model"
)

//Payment DTO for REST requests binding
type RestPayment struct {
	Accepted    bool
	PaymentType model.PaymentType
}