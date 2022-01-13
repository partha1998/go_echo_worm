package controllers

import (
	"fmt"
	"go_echo_gorm/db"
	"go_echo_gorm/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Sample(c echo.Context) error {
	return c.String(http.StatusOK, "Heelo from server")
}

func AddProduct(c echo.Context) error {

	// type Product struct {
	// 	Name        string `json:"Product_Name"`
	// 	Price       uint   `json:"Price"`
	// 	IsAvailable bool   `json:"Is_Available"`
	// }

	var reqBody models.Product

	err := c.Bind(&reqBody)
	if err != nil {
		return err
	}

	fmt.Println(reqBody)

	result := db.Database.Db.Create(&reqBody)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Data is not stored ")
	}
	return c.JSON(http.StatusOK, "Data stored successfully")
}

func GetAllProduct(c echo.Context) error {
	var products []models.Product
	db.Database.Db.Model(&models.Product{}).Find(&products)
	return c.JSON(http.StatusOK, products)
}
