package handlers

import (
	"context"
	"log"
	"mongo/db/dbface"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

// create mongo DB
func insertProduct(ctx context.Context, products []Product, collection dbface.Collection) ([]interface{}, error) {

	var insertIds []interface{}

	for _, product := range products {

		product.ID = primitive.NewObjectID()

		insertId, err := collection.InsertOne(ctx, product)
		if err != nil {
			log.Fatalf("inserOne error ->%s", err)
			return nil, err
		}

		insertIds = append(insertIds, insertId.InsertedID)

	}
	return insertIds, nil
}

func (h *ProductHandler) CreateProducts(c echo.Context) error {

	var products []Product


	if err := c.Bind(&products); err != nil {
		return err
	}
	
	Ids, err := insertProduct(context.Background(), products, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, Ids)
}

