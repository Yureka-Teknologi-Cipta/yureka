package yureka_mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func Connect(timeout time.Duration, URI, dbName string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOptions := options.Client()
	clientOptions.ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if dbName == "" {
		connString, _ := connstring.Parse(URI)
		dbName = connString.Database
	}

	return client.Database(dbName)
}
