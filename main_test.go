package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type mockCollection struct {
}

func (m *mockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{}
	return c, nil
}

func TestInsertData(t *testing.T) {
	mockCol := &mockCollection{}

	us := User{
		FirstName: "BORAT",
		LastName:  "ACUMVA",
	}
	res, err := insertData(mockCol, us)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}
