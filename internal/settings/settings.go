package settings

import (
	"os"
)

type Config struct {
	MongoURI          string
	RedisAddr         string
	RabbitMQURL       string
	RabbitMQQueueName string
}

func Load() *Config {
	return &Config{
		MongoURI:          os.Getenv("MONGO_URI"),
		RedisAddr:         os.Getenv("REDIS_ADDR"),
		RabbitMQURL:       os.Getenv("RABBITMQ_URL"),
		RabbitMQQueueName: "meli-challenge",
	}
}
