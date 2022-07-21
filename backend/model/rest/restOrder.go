package rest

//Order DTO for REST requests binding
type RestOrder struct {
	UserId     int
	Ammount    float32
	OrderItems []RestOrderItem
}
