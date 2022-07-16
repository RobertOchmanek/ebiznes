package rest

//Cart DTO for REST requests binding
type RestCart struct {
	UserId    int
	CartItems []RestCartItem
}