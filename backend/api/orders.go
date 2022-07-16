package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/RobertOchmanek/ebiznes_go/model/rest"
	"net/http"
	"github.com/labstack/echo/v4"
)

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
    db.Where("id = ?", id).Preload("Payment").Preload("OrderItems").Find(&order)

	return c.JSON(http.StatusOK, order)
}

func CreateOrder(c echo.Context) error {

	//Bind json from request to object
	restOrder := new(rest.RestOrder)
	c.Bind(restOrder)

	//Obtain current database connection
	db := database.DbManager()

	currentCart := model.Cart{}
	//Preload all cart's items and delete them
    db.Where("user_id = ?",  restOrder.UserId).Preload("CartItems").Find(&currentCart)
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

	//Convert payment from REST DTO to model object
	payment := model.Payment{}
	payment.Accepted = restOrder.Payment.Accepted
	payment.PaymentType = restOrder.Payment.PaymentType

	//Save new order, order ID is added by GORM
	newOrder := model.Order{}
	newOrder.OrderItems = orderItems
	newOrder.Payment = payment
	db.Create(&newOrder)

	//Update user's orders to save association between objects
	user := model.User{}
	db.Where("id = ?", restOrder.UserId).Find(&user)
	user.Orders = append(user.Orders, newOrder)
	db.Save(&user)

	return c.JSON(http.StatusOK, newOrder)
}