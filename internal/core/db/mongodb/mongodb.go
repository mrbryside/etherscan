package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	pocketsCollection = "pockets"
)

func NewMongoClient(mongoUrl string) *mongo.Client {
	client, err := connect(context.Background(), mongoUrl)
	if err != nil {
		panic(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err.Error())
	}

	return client
}

func connect(
	ctx context.Context,
	connectionString string,
) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	return client, err
}
