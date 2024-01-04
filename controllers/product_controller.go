package controllers

import (
	"mvc-be20/models"
	"mvc-be20/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductController(c echo.Context) error {
	newProduct := models.Product{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	// Set the UserID based on the authenticated user or any other logic
	// Example: newProduct.UserID = authenticatedUserID

	err := repositories.InsertProduct(newProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data. insert failed",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "insert success",
	})
}

func GetAllProductsController(c echo.Context) error {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error read data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    products,
	})
}

func GetProductByIdController(c echo.Context) error {
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	productData, err := repositories.GetProductById(idParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "error product not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    productData,
	})
}

func UpdateProductByIdController(c echo.Context) error {
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	var updatedProduct models.Product
	errBind := c.Bind(&updatedProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	err := repositories.UpdateProductById(idParam, updatedProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error update " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func DeleteProductByIdController(c echo.Context) error {
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	err := repositories.DeleteProductById(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error delete " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}
