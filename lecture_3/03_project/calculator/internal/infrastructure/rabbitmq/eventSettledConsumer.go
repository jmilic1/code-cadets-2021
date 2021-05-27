package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pkg/errors"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// EventSettledConsumer consumes settled events from the desired RabbitMQ queue.
type EventSettledConsumer struct {
	channel Channel
	config  ConsumerConfig
}

// NewEventSettledConsumer creates and returns a new EventSettledConsumer.
func NewEventSettledConsumer(channel Channel, config ConsumerConfig) (*EventSettledConsumer, error) {
	_, err := channel.QueueDeclare(
		config.Queue,
		config.DeclareDurable,
		config.DeclareAutoDelete,
		config.DeclareExclusive,
		config.DeclareNoWait,
		config.DeclareArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "event settled consumer initialization failed")
	}

	return &EventSettledConsumer{
		channel: channel,
		config:  config,
	}, nil
}

// Consume consumes messages until the context is cancelled. An error will be returned if consuming
// is not possible.
func (c *EventSettledConsumer) Consume(ctx context.Context) (<-chan models.EventSettled, error) {
	msgs, err := c.channel.Consume(
		c.config.Queue,
		c.config.ConsumerName,
		c.config.AutoAck,
		c.config.Exclusive,
		c.config.NoLocal,
		c.config.NoWait,
		c.config.Args,
	)
	if err != nil {
		return nil, errors.Wrap(err, "event settled consumer failed to consume messages")
	}

	events := make(chan models.EventSettled)
	go func() {
		defer close(events)
		for msg := range msgs {
			var event models.EventSettled
			err := json.Unmarshal(msg.Body, &event)
			if err != nil {
				log.Println("Failed to unmarshal event settled message", msg.Body)
			}

			// Once context is cancelled, stop consuming messages.
			select {
			case events <- event:
			case <-ctx.Done():
				return
			}
		}
	}()

	return events, nil
}
