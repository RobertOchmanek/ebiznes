package model

import (
	"github.com/jinzhu/gorm"
)

//User table for DB migration
type User struct {
	gorm.Model
	Username   string
	OauthId    int
	OauthToken string
	UserToken  string
	Cart       Cart
	Orders     []Order
}