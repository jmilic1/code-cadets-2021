package services

import (
	requests "github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/api/controllers/models"
	queueDto "github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/infrastructure/rabbitmq/models"
)

// BetService implements event related functions.
type BetService struct {
	betReceivedPublisher BetReceivedPublisher
	idGenerator          IdGenerator
}

// NewBetService creates a new instance of EventService.
func NewBetService(betReceivedPublisher BetReceivedPublisher, idGenerator IdGenerator) *BetService {
	return &BetService{
		betReceivedPublisher: betReceivedPublisher,
		idGenerator:          idGenerator,
	}
}

// SendBet sends received bet to the queue.
func (b *BetService) SendBet(request requests.BetRequestDto) error {
	id, err := b.idGenerator.GetRandomUUID()
	if err != nil {
		return err
	}

	betQueueDto := queueDto.BetQueueDto{
		Id:                   id,
		CustomerId:           request.CustomerId,
		SelectionId:          request.SelectionId,
		SelectionCoefficient: request.SelectionCoefficient,
		Payment:              request.Payment,
	}

	return b.betReceivedPublisher.Publish(betQueueDto)
}
