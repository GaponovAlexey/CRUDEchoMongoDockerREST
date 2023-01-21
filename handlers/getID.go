package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

func (h *ProductHandler) GetId(c echo.Context) error {
	var prod Product
	docId, err := primitive.ObjectIDFromHex(c.Param("id"))
	log.Println(docId)

	if err != nil {
		return err
	}
	filter := bson.M{"_id": docId}
	res := h.Col.FindOne(ctx, filter)

	if err := res.Decode(&prod); err != nil {
		return err
	}

	// res, err := h.Col.FindOne(ctx, filter)

	return c.JSON(http.StatusOK, prod)
}
