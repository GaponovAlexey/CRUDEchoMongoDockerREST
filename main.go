package main

import (
	"context"
	"log"
	"mongo/db/dbface"

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

func insertData(col dbface.CollectionAPI, user User) (*mongo.InsertOneResult, error) {
	res, err := col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main() {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	cancelError(err)

	db := client.Database("tronics").Collection("products")

	us := User{
		FirstName: "BORAT",
		LastName:  "BARBARA",
	}
	res, err := insertData(db, us)
	cancelError(err)

	log.Println(res)
}
