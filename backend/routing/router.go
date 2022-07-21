package routing

import (
	"github.com/RobertOchmanek/ebiznes_go/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ordersPath = "/orders"
var productsPath = "/products"
var productsPathWithId = "/products/:id"
var categoriesPath = "/categories"
var categoriesPathWithId = "/categories/:id"

func Init() *echo.Echo {

	//Create server instance
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	//Initialize endpoints and handler methods
	e.GET(ordersPath, api.GetOrders)
	e.GET("/orders/:id", api.GetOrder)
	e.POST(ordersPath, api.CreateOrder)

	e.GET(productsPath, api.GetProducts)
	e.GET(productsPathWithId, api.GetProduct)
	e.POST(productsPath, api.AddProduct)
	e.PUT(productsPathWithId, api.UpdateProduct)

	e.GET("/users", api.GetUsers)
	e.GET("/users/:userToken", api.GetUser)

	e.GET(categoriesPath, api.GetCategories)
	e.GET(categoriesPathWithId, api.GetCategory)
	e.POST(categoriesPath, api.AddCategory)
	e.PUT(categoriesPathWithId, api.UpdateCategory)
	e.DELETE(categoriesPathWithId, api.RemoveCategory)

	e.GET("/cartItems/:userId", api.GetCartItems)
	e.PUT("/cart", api.UpdateCart)

	e.GET("/oauth/login", api.OauthLoginUrl)
	e.GET("/oauth/logout/:userId", api.OauthLogoutUrl)
	e.GET("/oauth/callback", api.OauthCallback)

	e.POST("/payment", api.HandlePayment)

	return e
}
