package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/queue"
)

type RabbitMQProducer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQProducer(url string) (queue.Producer, error) {
	log.Printf("Connecting to RabbitMQ at %s", url)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	exchangeName := "meli-challenge"
	log.Printf("Declaring exchange %s", exchangeName)

	err = ch.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	queueName := "meli-shorten-url-queue"
	log.Printf("Declaring queue %s", queueName)

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	routingKey := "url.shorten"
	log.Printf("Binding queue %s to exchange %s with routing key %s", queueName, exchangeName, routingKey)

	err = ch.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &RabbitMQProducer{
		conn:    conn,
		channel: ch,
	}, nil
}

func (p *RabbitMQProducer) PublishShortenURL(ctx context.Context, shortURL *entity.ShortURL) error {
	body, err := json.Marshal(shortURL)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	exchangeName := "meli-challenge"
	routingKey := "url.shorten"

	log.Printf("Publishing message to exchange %s with routing key %s", exchangeName, routingKey)

	err = p.channel.PublishWithContext(
		ctx,
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Message published successfully")
	return nil
}
