package engine

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Handler interface {
	HandleBets(ctx context.Context, bets <-chan rabbitmqmodels.Bet)
	HandleEventsSettled(ctx context.Context, eventsSettled <-chan rabbitmqmodels.EventSettled) <-chan rabbitmqmodels.BetCalculated
}
