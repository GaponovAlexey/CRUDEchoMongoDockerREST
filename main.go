package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

)

var (
	e = echo.New()
)

func main() {

	e.GET("/", createProduct)
	e.Logger.Fatal(e.Start(":3000"))
}

func createProduct(e echo.Context) error {
	return e.JSON(http.StatusOK, "you")
}
