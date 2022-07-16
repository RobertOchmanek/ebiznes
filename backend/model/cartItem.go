package model

import (
	"github.com/jinzhu/gorm"
)

//CartItem table for DB migration
type CartItem struct {
	gorm.Model
	CartId    int
	ProductId int
	Quantity  int
}