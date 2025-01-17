package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/infrastructure/rabbitmq/models"
)

const contentTypeTextPlain = "text/plain"

// BetReceivedPublisher handles received bets queue publishing.
type BetReceivedPublisher struct {
	exchange  string
	queueName string
	mandatory bool
	immediate bool
	publisher QueuePublisher
}

// NewBetReceivedPublisher creates a new instance of BetReceivedPublisher.
func NewBetReceivedPublisher(
	exchange string,
	queueName string,
	mandatory bool,
	immediate bool,
	publisher QueuePublisher,
) *BetReceivedPublisher {
	return &BetReceivedPublisher{
		exchange:  exchange,
		queueName: queueName,
		mandatory: mandatory,
		immediate: immediate,
		publisher: publisher,
	}
}

// Publish publishes a received bet message to the queue.
func (p *BetReceivedPublisher) Publish(bet models.BetQueueDto) error {
	betReceivedJson, err := json.Marshal(bet)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		p.exchange,
		p.queueName,
		p.mandatory,
		p.immediate,
		amqp.Publishing{
			ContentType: contentTypeTextPlain,
			Body:        betReceivedJson,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent %s", betReceivedJson)
	return nil
}
