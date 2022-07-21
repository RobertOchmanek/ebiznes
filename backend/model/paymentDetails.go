package model

import "github.com/jinzhu/gorm"

//Payment details table for DB migration
type PaymentDetails struct {
	gorm.Model
	Amount      int64
	PayerEmail  string
	Description string
}
