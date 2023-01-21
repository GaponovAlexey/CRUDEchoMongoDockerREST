package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	docId, err := primitive.ObjectIDFromHex(c.Param("id"))
	log.Println(docId)

	if err != nil {
		return err
	}
	filter := bson.M{"_id": docId}
	res, err := h.Col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	
	return c.JSON(http.StatusOK, res)
}
