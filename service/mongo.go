package service

import (
	"fmt"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewMongoConnection(host string, port int) (*mongo.Client, error) {

	
	client, err := mongo.Connect(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CheckIfConnected(client *mongo.Client) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	return client.Ping(ctx, readpref.Primary()) == nil
}