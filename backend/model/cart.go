package model

import (
	"github.com/jinzhu/gorm"
)

//Cart table for DB migration
type Cart struct {
	gorm.Model
	UserId    int
	CartItems []CartItem
}