package model

import (
	"github.com/jinzhu/gorm"
)

//OrderItem table for DB migration
type OrderItem struct {
	gorm.Model
	OrderId   int
	ProductId int
	Quantity  int
}