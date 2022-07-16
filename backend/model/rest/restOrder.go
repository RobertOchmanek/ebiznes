package rest

//Order DTO for REST requests binding
type RestOrder struct {
	UserId     int
	Payment    RestPayment
	OrderItems []RestOrderItem
}