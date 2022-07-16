package model

import (
	"github.com/jinzhu/gorm"
)

//Payment table for DB migration
type Payment struct {
	gorm.Model
	OrderId     int
	Accepted    bool
	PaymentType PaymentType
}