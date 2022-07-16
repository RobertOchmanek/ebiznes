package model

import (
	"github.com/jinzhu/gorm"
)

//Product table for DB migration
type Product struct {
	gorm.Model
	CategoryId int
	Name       string
	Price      float32
	Image	   string
}