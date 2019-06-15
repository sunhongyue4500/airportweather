package db

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func TestMongo(t *testing.T) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	t.Log(client, err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 做超时控制
	err = client.Connect(ctx)
	if err != nil {
		t.Fatalf(err.Error())
	}
	collection := client.Database("CHUAV").Collection("airportweather")

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		t.Fatalf(err.Error())
	}
	id := res.InsertedID
	t.Log(id)
}
