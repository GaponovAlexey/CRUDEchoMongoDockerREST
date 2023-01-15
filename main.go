package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type User struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}

var (
	ctxB = context.WithTimeout(context.)
)

func cancelError(e error) {
	if e != nil {
		log.Println(e)
	}
}

func insertData(col *mongo.Collection, user User) (*mongo.InsertOneResult, error) {
	res, err := col.InsertOne(ctxB, user)
}

func main() {
	fmt.Println("start")
	// ctx, cancel := context.WithTimeout(ctxB, 2*time.Second)
	// defer cancel()
	client, err := mongo.Connect(ctxB, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer cancelError(err)
	db := client.Database("tronics").Collection("products")

	update := bson.M{"product_name": "Iphone33"}

	ok, err := db.DeleteOne(ctx, update)

	log.Println(ok)

	//end
	fmt.Println("end")
}
