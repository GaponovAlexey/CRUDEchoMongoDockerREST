package handlers

import (
	"context"
	"fmt"
	"mongo/db/dbface"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

)

// GET
func findProducts(ctx context.Context, collection dbface.Collection) ([]Product, error) {
	var product []Product
	
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Errorf("find Product Error")
		return nil, err
	}
	err = cursor.All(ctx, &product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (h *ProductHandler) GetProduct(c echo.Context) error {

	products, err := findProducts(context.Background(), h.Col)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}