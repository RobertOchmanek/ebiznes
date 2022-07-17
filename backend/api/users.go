package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/RobertOchmanek/ebiznes_go/model/rest"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUsers(c echo.Context) error {

	//Obtain current database connection and fetch users
	db := database.DbManager()
	users := []model.User{}
	db.Find(&users)

	//Return REST DTOs to hide tokens
	restUsers := []rest.RestUser{}
	for _, user := range users {
		restUser := rest.RestUser{}
		restUser.ID = int(user.ID)
		restUser.Username = user.Username
		restUser.OauthId = user.OauthId
		restUser.Cart = user.Cart
		restUser.Orders = user.Orders
		restUsers = append(restUsers, restUser)
	}

	return c.JSON(http.StatusOK, restUsers)
}

func GetUser(c echo.Context) error {

	//Get user ID from query param
	userToken := c.Param("userToken")

	//Obtain current database connection and fetch user by ID
	db := database.DbManager()
	user := model.User{}
	//Preload user's orders and cart and include in response
	db.Where("user_token = ?", userToken).Preload("Orders.OrderItems").Preload("Orders.Payment").Preload("Cart.CartItems").Find(&user)

	//Return REST DTO to hide tokens
	restUser := rest.RestUser{}
	restUser.ID = int(user.ID)
	restUser.Username = user.Username
	restUser.OauthId = user.OauthId
	restUser.Cart = user.Cart
	restUser.Orders = user.Orders

	return c.JSON(http.StatusOK, restUser)
}

func UserExists(Username string) bool {

	//Obtain current database connection and fetch user by username
	db := database.DbManager()
	user := model.User{}
	db.Where("username = ?", Username).Find(&user)

	//User exists if object returned from DB does not contain empty fields
	return user.Username != ""
}
