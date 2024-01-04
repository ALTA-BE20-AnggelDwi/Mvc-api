package main

import (
	"fmt"
	"mvc-be20/config"
	"mvc-be20/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Printf("running")
	config.InitDB()
	config.InitialMigration()

	e := echo.New()
	routes.InitRoute(e)

	//start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
