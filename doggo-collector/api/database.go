package main

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	//TODO add card Create Read and Deletes
}

type mongoDatabase struct {
	client *mongo.Client
}

func NewMongoDatabase(client *mongo.Client) Database {
	return &mongoDatabase{client: client}
}
