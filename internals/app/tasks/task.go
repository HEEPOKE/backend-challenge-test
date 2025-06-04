package tasks

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LogUserCountTask(conn *mongo.Client) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	collection := conn.Database("test").Collection("users")

	go func() {
		for {
			select {
			case <-ticker.C:
				count, err := collection.CountDocuments(context.Background(), bson.M{})
				if err != nil {
					log.Printf("Error counting users: %v", err)
					continue
				}
				log.Printf("Current user count: %d", count)
			}
		}
	}()
}
