package api

import (
	"net/http"
	"strconv"

	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/RobertOchmanek/ebiznes_go/model/rest"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

const stripeApiKey = ""

func GetOrders(c echo.Context) error {

	//Obtain current database connection and fetch orders
	db := database.DbManager()
	orders := []model.Order{}

	db.Preload("Payment").Preload("OrderItems").Find(&orders)

	return c.JSON(http.StatusOK, orders)
}

func GetOrder(c echo.Context) error {

	//Get order ID from query param
	id := c.Param("id")

	//Obtain current database connection and fetch order by ID
	db := database.DbManager()
	order := model.Order{}
	//Preload all order's items and payment and include in response
	db.Where(idEquals, id).Preload("Payment").Preload("OrderItems").Find(&order)

	return c.JSON(http.StatusOK, order)
}

func CreateOrder(c echo.Context) error {

	//Bind json from request to object
	restOrder := new(rest.RestOrder)
	c.Bind(restOrder)

	//Obtain current database connection
	db := database.DbManager()

	//Obtain user from DB
	user := model.User{}
	db.Where(idEquals, restOrder.UserId).Find(&user)

	currentCart := model.Cart{}
	//Preload all cart's items and delete them
	db.Where("user_id = ?", restOrder.UserId).Preload("CartItems").Find(&currentCart)
	for _, cartItem := range currentCart.CartItems {
		db.Delete(&cartItem)
	}

	//Create order's items
	orderItems := []model.OrderItem{}
	for _, orderItem := range restOrder.OrderItems {
		item := model.OrderItem{}
		item.ProductId = orderItem.ProductId
		item.Quantity = orderItem.Quantity
		orderItems = append(orderItems, item)
	}

	stripe.Key = stripeApiKey

	//NOTE: multiplied by 100 as Stripe accepts ammount in cents
	ammount := int(restOrder.Ammount * 100)
	description := "Ordered items: "
	accepted := true

	for _, orderItem := range orderItems {
		product := model.Product{}
		db.Where(idEquals, orderItem.ProductId).Find(&product)
		description += (strconv.Itoa(orderItem.Quantity) + "x " + product.Name + ", ")
	}

	//NOTE: remove last ", " from description
	description = description[:(len(description) - 2)]

	//TODO: payment alert based on accepted
	//TODO: get real email from oauth
	//TODO: add card number, CCV and expiring
	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(int64(ammount)),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String(description),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String("temp@emial.com")})

	if err != nil {
		accepted = false
	}

	//Create model object for payment
	payment := model.Payment{}
	payment.Accepted = accepted
	payment.PaymentType = model.CreditCard
	payment.Amount = restOrder.Ammount
	payment.Description = description

	//Save new order, order ID is added by GORM
	newOrder := model.Order{}
	newOrder.OrderItems = orderItems
	newOrder.Payment = payment
	db.Create(&newOrder)

	//Update user's orders to save association between objects
	user.Orders = append(user.Orders, newOrder)
	db.Save(&user)

	return c.JSON(http.StatusOK, accepted)
}
