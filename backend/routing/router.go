package routing

import (
	"github.com/RobertOchmanek/ebiznes_go/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {

	//Create server instance
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	  })) 

	//Initialize endpoints and handler methods
	e.GET("/orders", api.GetOrders)
	e.GET("/orders/:id", api.GetOrder)
	e.POST("/orders", api.CreateOrder)

	e.GET("/products", api.GetProducts)
	e.GET("/products/:id", api.GetProduct)
	e.POST("/products", api.AddProduct)
	e.PUT("/products/:id", api.UpdateProduct)

	e.GET("/users", api.GetUsers)
	e.GET("/users/:id", api.GetUser)
	e.POST("/users", api.AddUser)
	e.PUT("/users/:id", api.UpdateUser)

	e.GET("/categories", api.GetCategories)
	e.GET("/categories/:id", api.GetCategory)
	e.POST("/categories", api.AddCategory)
	e.PUT("/categories/:id", api.UpdateCategory)
	e.DELETE("/categories/:id", api.RemoveCategory)

	e.GET("/cartItems/:userId", api.GetCartItems)
	e.PUT("/cart", api.UpdateCart)

	return e
}