package queue

import (
	"context"
)

// Consumer define a interface para consumir mensagens do RabbitMQ
type Consumer interface {
	// Start inicia o consumo de mensagens
	Start(ctx context.Context) error
	// Stop para o consumo de mensagens
	Stop() error
}
