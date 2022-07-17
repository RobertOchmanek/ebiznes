package rest

import (
	"github.com/RobertOchmanek/ebiznes_go/model"
)

//User DTO for REST requests binding
type RestUser struct {
	ID       int
	Username string
	OauthId  int
	Cart     model.Cart
	Orders   []model.Order
}