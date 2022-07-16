package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {

	//Obtain current database connection and fetch products
	db := database.DbManager()
	products := []model.Product{}
	db.Find(&products)

	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {

	//Get product ID from query param
	id := c.Param("id")

	//Obtain current database connection and fetch product by ID
	db := database.DbManager()
	product := model.Product{}
    db.Where("id = ?", id).Find(&product)

	return c.JSON(http.StatusOK, product)
}

func AddProduct(c echo.Context) error {

	//Bind json from request to object
	newProduct := new(model.Product)
	c.Bind(newProduct)

	//Obtain current database connection and save new product
	db := database.DbManager()
	db.Create(&newProduct)

	return c.JSON(http.StatusOK, newProduct)
}

func UpdateProduct(c echo.Context) error {

	//Get product ID from query param
	id := c.Param("id")

	//Bind json from request to object
	updatedProduct := new(model.Product)
	c.Bind(updatedProduct)

	//Obtain current database connection and update product by ID
	db := database.DbManager()
	product := model.Product{}
    db.Where("id = ?", id).Find(&product)

	//Update and save DB object
	product.CategoryId = updatedProduct.CategoryId
	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.Image = updatedProduct.Image
	db.Save(&product)

	return c.JSON(http.StatusOK, product)
}