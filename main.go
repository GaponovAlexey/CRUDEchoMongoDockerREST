package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type User struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}

var (
	ctx = context.Background()
)

func cancelError(e error) {
	if e != nil {
		log.Println(e)
	}
}

func insertData(col *mongo.Collection, user User) (*mongo.InsertOneResult, error) {
	res, err := col.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error insert", err)
	}
	return res, nil
}

func main() {
	fmt.Println("start")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	cancelError(err)

	db := client.Database("tronics").Collection("products")

	// update := bson.M{"product_name": "Iphone33"}
	us := User{
		FirstName: "BORAT",
		LastName:  "ACUMVA",
	}
	res, err := insertData(db, us)
	cancelError(err)


	
	log.Println(res)

	//end
	fmt.Println("end")
}
