package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"mongo/db/dbface"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

func modifyProduct(ctx context.Context, id string, reqBody io.ReadCloser, col dbface.Collection) (Product, error) {

	var product Product

	gId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": gId}

	res := col.FindOne(ctx, filter)

	if err := res.Decode(&product); err != nil {
		return product, err
	}

	if err := json.NewDecoder(reqBody).Decode(&product); err != nil {
		return product, err
	}

	_, err := col.UpdateOne(ctx, filter, bson.M{"$set": product})
	if err != nil {
		return product, err
	}

	return product, nil

}



// POST
func (h *ProductHandler) PutProduct(c echo.Context) error {
	id := c.Param("id")

	prod, err := modifyProduct(ctx, id, c.Request().Body, h.Col)
	if err != nil {
		log.Fatal("ModufyProduct Error", err)
		return err
	}

	return c.JSON(http.StatusOK, prod)
}
