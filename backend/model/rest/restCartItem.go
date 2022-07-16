package rest

import (
	"time"
)

//Cart item DTO for REST requests binding
type RestCartItem struct {
	ID         int
	CategoryId int
	Name       string
	Price      float32
	Image	   string
	Quantity   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}