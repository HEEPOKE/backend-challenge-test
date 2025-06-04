package main

import (
	"log"

	"github.com/HEEPOKE/backend-challenge-test/internals/app/tasks"
	server "github.com/HEEPOKE/backend-challenge-test/internals/servers"
	"github.com/HEEPOKE/backend-challenge-test/pkg/configs"
	"github.com/HEEPOKE/backend-challenge-test/pkg/databases"
)

func main() {
	_, err := configs.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}

	client, err := databases.ConnectMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	tasks.LogUserCountTask(client)

	route := server.NewServer(client)
	fib := route.Init()

	err = fib.Listen("3000")
	if err != nil {
		log.Fatalf("Failed To Start The Server: %v", err)
	}
}
