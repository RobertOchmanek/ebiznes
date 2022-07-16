package model

import (
	"github.com/jinzhu/gorm"
)

//Category table for DB migration
type Category struct {
	gorm.Model
	Name     string
	Products []Product
}