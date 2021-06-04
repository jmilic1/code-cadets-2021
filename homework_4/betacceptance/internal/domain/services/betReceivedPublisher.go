package services

import "github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/infrastructure/rabbitmq/models"

// BetReceivedPublisher handles event update queue publishing.
type BetReceivedPublisher interface {
	Publish(bet models.BetQueueDto) error
}
