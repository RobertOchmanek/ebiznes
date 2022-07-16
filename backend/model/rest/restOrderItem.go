package rest

//Order item DTO for REST requests binding
type RestOrderItem struct {
	ProductId int
	Quantity  int
}