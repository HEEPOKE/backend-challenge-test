package databases

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/backend-challenge-test/pkg/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(configs.Cfg.MONGODB_URI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to MongoDB")

	return client, nil
}
