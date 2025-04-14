package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"github.com/rbcorrea/meli-challenge/internal/infrastructure/queue"
	"github.com/rbcorrea/meli-challenge/internal/infrastructure/repository/mongo"
	"github.com/redis/go-redis/v9"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGODB_URL")

	client, err := mongodriver.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("meli-challenge").Collection("short_urls")
	var repo repository.ShortenURLRepository = mongo.NewMongoShortenURLRepository(collection)

	redisURL := os.Getenv("REDIS_URL")

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	rabbitURL := os.Getenv("RABBITMQ_URL")

	consumer, err := queue.NewConsumer(rabbitURL, repo, redisClient)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	if err := consumer.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	select {}
}
