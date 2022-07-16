package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/RobertOchmanek/ebiznes_go/model/rest"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetCartItems(c echo.Context) error {

	//Get user ID from query param
	userId := c.Param("userId")

	//Obtain current database connection and fetch cart by user ID
	db := database.DbManager()
	cart := model.Cart{}
	//Preload all cart's items to obtain quantity and ID of each item
    db.Where("user_id = ?", userId).Preload("CartItems").Find(&cart)

	//For each cart item fetch product by ID and convert it to REST DTO
	restCartItems := []rest.RestCartItem{}
	for _, cartItem := range cart.CartItems {

		fetchedProduct := model.Product{}
		restCartItem := rest.RestCartItem{}

		db.Where("id = ?", cartItem.ProductId).Find(&fetchedProduct)

		restCartItem.ID = int(fetchedProduct.ID)
		restCartItem.CategoryId = fetchedProduct.CategoryId
		restCartItem.Name = fetchedProduct.Name
		restCartItem.Price = fetchedProduct.Price
		restCartItem.Image = fetchedProduct.Image
		restCartItem.Quantity = cartItem.Quantity
		restCartItem.CreatedAt = fetchedProduct.CreatedAt
		restCartItem.UpdatedAt = fetchedProduct.UpdatedAt
		restCartItem.DeletedAt = fetchedProduct.DeletedAt

		restCartItems = append(restCartItems, restCartItem)
	}

	return c.JSON(http.StatusOK, restCartItems)
}

func UpdateCart(c echo.Context) error {

	//Bind json from request to object
	updatedCart := new(rest.RestCart)
	c.Bind(updatedCart)

	//Obtain current database connection
	db := database.DbManager()

	currentCart := model.Cart{}
	//Preload all cart's items and delete them
    db.Where("user_id = ?", updatedCart.UserId).Preload("CartItems").Find(&currentCart)

	for _, cartItem := range currentCart.CartItems {
		db.Delete(&cartItem)
	}

	//Re-create cart's items with updated values
	updatedCartItems := []model.CartItem{}
	for _, cartItem := range updatedCart.CartItems {
		updatedCartItem := model.CartItem{}
		updatedCartItem.CartId = int(currentCart.ID)
		updatedCartItem.ProductId = cartItem.ID
		updatedCartItem.Quantity = cartItem.Quantity
		updatedCartItems = append(updatedCartItems, updatedCartItem)
	}

	currentCart.CartItems = updatedCartItems

	//Save updated cart as current one
	db.Save(&currentCart)
	return c.JSON(http.StatusOK, currentCart)
}