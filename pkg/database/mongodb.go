package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// MongoDB is struct to connect MongoDB
type MongoDB struct {
	URI string
}

func NewMongoDB(uri string) *MongoDB {
	return &MongoDB{
		URI: uri,
	}
}

// Connect is method to connect MongoDB
func (m *MongoDB) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")) // todo replace to config.mongoURL
	if err != nil {
		log.Fatal()
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal()
	}

	return nil
}

// Disconnect is method to disconnect MongoDB
func (m *MongoDB) Disconnect() error {
	return nil
}
