package handlers

import (
	"context"
	"fmt"
	"mongo/db/dbface"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

var (
	v = validator.New()
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

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

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

// create mongo DB
func (h *ProductHandler) CreateProducts(c echo.Context) error {

	var products []Product

	c.Echo().Validator = &ProductValidator{validator: v}

	if err := c.Bind(&products); err != nil {
		return err
	}
	for _, product := range products {
		err := c.Validate(product)
		if err != nil {
			log.Fatal("Unable to validate the product <---", err)
			return err
		}

	}
	Ids, err := insertProduct(context.Background(), products, h.Col)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Ids)
}
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
