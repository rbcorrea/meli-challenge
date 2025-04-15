package main

import (
	"context"
	"log"
	"os"

	"github.com/rbcorrea/meli-challenge/internal/application/usecase"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"github.com/rbcorrea/meli-challenge/internal/infrastructure/api"
	"github.com/rbcorrea/meli-challenge/internal/infrastructure/queue"
	"github.com/rbcorrea/meli-challenge/internal/infrastructure/repository/mongo"
	"github.com/redis/go-redis/v9"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var ctx = context.Background()

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

	producer, err := queue.NewRabbitMQProducer(rabbitURL)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ producer: %v", err)
	}

	shortenUseCase := usecase.NewShortenURLUseCase(repo, producer, redisClient)
	searchByCodeUseCase := usecase.NewSearchByCodeUseCase(repo, redisClient)
	redirectUseCase := usecase.NewRedirectUseCase(repo, redisClient)
	deleteUseCase := usecase.NewDeleteURLUseCase(repo, redisClient)

	app := api.NewApp(shortenUseCase, searchByCodeUseCase, redirectUseCase, deleteUseCase)

	log.Fatal(app.Listen(":8080"))
}
