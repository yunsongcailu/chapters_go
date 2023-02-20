package store

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var M *MongoStore

type MongoStore struct {
	Client *mongo.Client
	DB     *mongo.Database
	Coll   map[string]*mongo.Collection
}

func NewMongoStore(ctx context.Context) error {
	user := "root"
	pass := "root"
	addr := "127.0.0.1"
	port := 27017
	db := "car"
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pass, addr, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	M = &MongoStore{
		Client: client,
		DB:     nil,
		Coll:   nil,
	}
}
