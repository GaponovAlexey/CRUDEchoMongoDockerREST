package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

)

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	fId := c.Param("id")
	filter := bson.M{"_id": fId}
	_, err := h.Col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}
