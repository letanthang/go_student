package db

import (
	"context"
	"fmt"
	"time"

	"github.com/letanthang/go_student/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

const (
	DBName = "go0110"
	Col    = "student"
)

func init() {
	uri := config.Config.Mongo.Uri
	fmt.Println("Connect MongoDb with uri", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	Client = client
}
