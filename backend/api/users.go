package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {

	//Obtain current database connection and fetch users
	db := database.DbManager()
	users := []model.User{}
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {

	//Get user ID from query param
	id := c.Param("id")

	//Obtain current database connection and fetch user by ID
	db := database.DbManager()
	user := model.User{}
	//Preload user's orders and cart and include in response
    db.Where("id = ?", id).Preload("Orders.OrderItems").Preload("Orders.Payment").Preload("Cart.CartItems").Find(&user)

	return c.JSON(http.StatusOK, user)
}

func UserExists(Username string) bool {

	//Obtain current database connection and fetch user by username
	db := database.DbManager()
	user := model.User{}
	db.Where("username = ?", Username).Find(&user)

	//User exists if object returned from DB does not contain empty fields
	if (user.Username == "") {
		return false
	}

	return true
}