package consumer

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Consumer offers methods for consuming from input queues.
type Consumer struct {
	betConsumer          BetConsumer
	eventSettledConsumer EventSettledConsumer
}

// New creates and returns a new Consumer.
func New(betConsumer BetConsumer, eventSettledConsumer EventSettledConsumer) *Consumer {
	return &Consumer{
		betConsumer:          betConsumer,
		eventSettledConsumer: eventSettledConsumer,
	}
}

// ConsumeBets consumes bets queue.
func (c *Consumer) ConsumeBets(ctx context.Context) (<-chan rabbitmqmodels.Bet, error) {
	return c.betConsumer.Consume(ctx)
}

// ConsumeEventsSettled consumes bets calculated queue.
func (c *Consumer) ConsumeEventsSettled(ctx context.Context) (<-chan rabbitmqmodels.EventSettled, error) {
	return c.eventSettledConsumer.Consume(ctx)
}
