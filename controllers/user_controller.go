package controllers

import (
	"fmt"
	"mvc-be20/entities"
	"mvc-be20/models"
	"mvc-be20/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// insert data user
func CreateUserController(c echo.Context) error {
	newUser := entities.UserCore{}
	errBind := c.Bind(&newUser) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	// simpan ke DB
	errInsert := repositories.InsertUser(newUser)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data. insert failed",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "insert success",
	})
}

// get user by id
func GetUserByIdController(c echo.Context) error {
	userId := c.Param("user_id")
	userIdParam, errConv := strconv.Atoi(userId)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. user_id should be number",
		})
	}

	userData, err := repositories.GetUserById(userIdParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error get user " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    userData,
	})
}

// read data user
func GetAllUserController(c echo.Context) error {

	results, errSelect := repositories.SelectAllUsers()
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error read data. " + errSelect.Error(),
		})
	}
	fmt.Println("users:", results)
	// proses mapping dari core ke response
	var usersResult []entities.UserResponse
	for _, value := range results {
		usersResult = append(usersResult, entities.UserResponse{
			ID:    value.ID,
			Name:  value.Name,
			Email: value.Email,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    usersResult,
	})
}

func UpdateUserByIdController(c echo.Context) error {
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}
	var userData = models.User{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}
	err := repositories.UpdateUserById(idParam, userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	err := repositories.DeleteUserById(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error delete " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

// get user product
func GetProductsByUserIdController(c echo.Context) error {
	userId := c.Param("user_id")
	userIdParam, errConv := strconv.Atoi(userId)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. user_id should be number",
		})
	}

	products, err := repositories.GetProductsByUserId(userIdParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error get products " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    products,
	})
}
