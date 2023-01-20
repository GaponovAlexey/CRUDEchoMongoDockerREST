package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

)

func (h *ProductHandler) PutProduct(c echo.Context) error {
 


	return c.JSON(http.StatusOK, "you")
}
