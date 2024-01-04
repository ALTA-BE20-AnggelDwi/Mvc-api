package routes

import (
	"mvc-be20/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	// create a new echo instance

	// define routes/ endpoint
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users", controllers.GetAllUserController)
	e.GET("/users/:user_id", controllers.GetUserByIdController)
	e.PUT("/users/:user_id", controllers.UpdateUserByIdController)
	e.DELETE("/users/:user_id", controllers.DeleteUserController)
	e.GET("/users/:user_id/products", controllers.GetProductsByUserIdController)

	e.POST("/products", controllers.CreateProductController)
	e.GET("/products", controllers.GetAllProductsController)
	e.GET("/products/:product_id", controllers.GetProductByIdController)
	e.PUT("/products/:product_id", controllers.UpdateProductByIdController)
	e.DELETE("/products/:product_id", controllers.DeleteProductByIdController)

	// return e
}
