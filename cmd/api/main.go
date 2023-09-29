package main

import (
	"C0lliNN/datadog-demo/internal"
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT")) 	
	if port == 0 {
		log.Fatal("port not provided")
	}

	mongouri := os.Getenv("MONGO_URI")
	if mongouri == "" {
		log.Fatal("mongo uri not provided")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongouri))
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("orders")
	repo := internal.NewOrderRepository(database)
	service := internal.NewOrderService(repo)
	server := internal.NewServer(service, port)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}