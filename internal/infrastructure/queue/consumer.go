package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type Consumer struct {
	conn       *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
	repository repository.ShortenURLRepository
	redis      *redis.Client
	done       chan struct{}
	wg         sync.WaitGroup
}

func NewConsumer(url string, repository repository.ShortenURLRepository, redis *redis.Client) (*Consumer, error) {
	log.Printf("Connecting to RabbitMQ at %s", url)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	exchangeName := "meli-challenge"
	log.Printf("Declaring exchange %s", exchangeName)

	err = channel.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	queueName := "meli-shorten-url-queue"
	log.Printf("Declaring queue %s", queueName)

	queue, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	routingKey := "url.shorten"
	log.Printf("Binding queue %s to exchange %s with routing key %s", queueName, exchangeName, routingKey)

	err = channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &Consumer{
		conn:       conn,
		channel:    channel,
		queue:      &queue,
		repository: repository,
		redis:      redis,
		done:       make(chan struct{}),
	}, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	msgs, err := c.channel.Consume(
		c.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-c.done:
				return
			case msg := <-msgs:
				var shortURL entity.ShortURL
				if err := json.Unmarshal(msg.Body, &shortURL); err != nil {
					fmt.Printf("Failed to unmarshal message: %v\n", err)
					continue
				}

				data, _ := json.Marshal(shortURL)
				if err := c.redis.Set(ctx, shortURL.Code, data, 24*time.Hour).Err(); err != nil {
					fmt.Printf("Failed to save to Redis: %v\n", err)
				}

				if err := c.repository.Save(ctx, &shortURL); err != nil {
					fmt.Printf("Failed to save to MongoDB: %v\n", err)
				}
			}
		}
	}()

	return nil
}

func (c *Consumer) Stop() error {
	close(c.done)
	c.wg.Wait()

	if err := c.channel.Close(); err != nil {
		return fmt.Errorf("failed to close channel: %w", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}
