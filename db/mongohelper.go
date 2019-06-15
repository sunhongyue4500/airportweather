package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

// 数据库名称
const databaseName = "CHUAV"

// 集合名称
const collectionName = "airportweathers"

var collection *mongo.Collection

// 连接数据库
func ConnMongoDB(ip string, port int) error {
	if mongoClient != nil {
		return nil
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", ip, port)))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 做超时控制
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	mongoClient = client
	collection = mongoClient.Database(databaseName).Collection(collectionName)
	return nil
}

func InsertManyDocs(documents []interface{}) (*mongo.InsertManyResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 做超时控制
	result, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}
	log.Printf("插入%d条数据\n", len(result.InsertedIDs))
	return result, nil
}

func DropCollection() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 做超时控制
	if err := collection.Drop(ctx); err != nil {
		return err
	}
	log.Println("删除Collection成功")
	return nil
}
