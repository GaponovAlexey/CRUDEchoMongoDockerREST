package handlers

import (
	"context"
	"log"
	"mongo/db/dbface"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    string             `json:"quantity" bson:"quantity"`
	Discount    int                `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor      string             `json:"vendor" bson:"vendor"`
	Accessories []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuID       string             `json:"sku_id" bson:"sku_id"`
}

type ProductHandler struct {
	Col dbface.Collection
}

func createProduct(ctx context.Context, products []Product, collection dbface.Collection) ([]interface{}, error) {
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

// create mongo DB
func (h *ProductHandler) CreateProducts(c echo.Context) error {

	var products []Product

	if err := c.Bind(&products); err != nil {
		return err
	}
	Ids, err := createProduct(context.Background(), products, h.Col)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Ids)
}
