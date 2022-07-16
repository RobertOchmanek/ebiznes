package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/RobertOchmanek/ebiznes_go/model/rest"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetCategories(c echo.Context) error {

	//Obtain current database connection and fetch categories
	db := database.DbManager()
	categories := []model.Category{}
	//Products in each category are excluded when getting all categories
	db.Find(&categories)

	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {

	//Get category ID from query param
	id := c.Param("id")

	//Obtain current database connection and fetch category by ID
	db := database.DbManager()
	category := model.Category{}
	//Preload all category's products and include in response
    db.Where("id = ?", id).Preload("Products").Find(&category)

	return c.JSON(http.StatusOK, category)
}

func AddCategory(c echo.Context) error {

	//Bind json from request to object
	restCategory := new(rest.RestCategory)
	c.Bind(restCategory)

	//Create and save new DB object
	newCategory := model.Category{}
	newCategory.Name = restCategory.Name

	//Obtain current database connection and save new category
	db := database.DbManager()
	db.Create(&newCategory)

	return c.JSON(http.StatusOK, newCategory)
}

func UpdateCategory(c echo.Context) error {

	//Get category ID from query param
	id := c.Param("id")

	//Bind json from request to object
	restCategory := new(rest.RestCategory)
	c.Bind(restCategory)

	//Obtain current database connection and update category by ID
	db := database.DbManager()
	category := model.Category{}
    db.Where("id = ?", id).Find(&category)

	//Update and save DB object
	category.Name = restCategory.Name
	db.Save(&category)

	return c.JSON(http.StatusOK, category)
}

func RemoveCategory(c echo.Context) error {

	//Get category ID from query param
	id := c.Param("id")

	//Obtain current database connection and remove category by ID
	db := database.DbManager()
	category := model.Category{}
	//Preload all category's products to check if it can be removed
    db.Where("id = ?", id).Preload("Products").Find(&category)

	if len(category.Products) > 0 {
		return c.JSON(http.StatusMethodNotAllowed, "Category can not be removed unless it does not contain any products")
	} else {
		db.Delete(&category)
		return c.JSON(http.StatusOK, category)
	}
}