package handlers

import (
	"context"
	"fmt"
	"mongo/db/dbface"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

)

// get id
func findProductsID(ctx context.Context, q url.Values, collection dbface.Collection) ([]Product, error) {
	var product []Product
	filter := make(map[string]interface{})
	for k, v := range q {
		filter[k] = v[0]
	}
	cursor, err := collection.Find(ctx, bson.M(filter))
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
func (h *ProductHandler) GetProductID(c echo.Context) error {

	products, err := findProductsID(context.Background(), c.QueryParams(), h.Col)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}
