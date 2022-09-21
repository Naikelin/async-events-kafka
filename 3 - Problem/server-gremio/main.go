package main

import (
	"guild/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	/* Routes */
	e.POST("/register", controllers.Register)

	e.Logger.Fatal(e.Start(":1323"))
}
