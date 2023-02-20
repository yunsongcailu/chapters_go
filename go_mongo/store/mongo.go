package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var M *MongoDB

// NewMongoDB uri := "mongodb://root:123456@localhost:27017"
func NewMongoDB(uri, dbName string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()
	MongoClient = client
	M = &MongoDB{
		DB:   client.Database(dbName),
		Coll: make(map[string]*mongo.Collection),
	}
	return nil
}
