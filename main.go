package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    string             `json:"quantity" bson:" quantity"`
	Discount    int                `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor      string             `json:"vendor" bson:"vendor"`
	Accessories []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuID       string             `json:"sku_id" bson:"sku_id"`
}

var iphone10 = Product{
	ID:          primitive.NewObjectID(),
	Name:        "Iphone15",
	Price:       900,
	Currency:    "CAD",
	Quantity:    "40",
	Vendor:      "apple",
	Accessories: []string{"charger", "headset", "slot"},
	SkuID:       "1234",
}
var trimmer = Product{
	ID:          primitive.NewObjectID(),
	Name:        "easy trimmer",
	Price:       100,
	Currency:    "CAD",
	Quantity:    "140",
	Vendor:      "Bosh",
	Accessories: []string{"charger", "headset", "slot"},
	SkuID:       "1234",
}

func cancelError(e error) {
	if e != nil {
		log.Println(e)
	}
}

func main() {
	log.Println("start")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	cancelError(err)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	cancelError(err)

	db := client.Database("tronics")
	collection := db.Collection("products")
	res, err := collection.InsertOne(context.Background(), iphone10)
	cancelError(err)
	log.Println(res.InsertedID.(primitive.ObjectID).Timestamp())

}
