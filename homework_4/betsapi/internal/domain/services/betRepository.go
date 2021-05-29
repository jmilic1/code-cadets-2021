package services

import (
	"context"

	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/domain/models"
)

type BetRepository interface {
	GetBetById(ctx context.Context, id string) (domainmodels.Bet, bool, error)
	GetBetsByCustomerId(ctx context.Context, customerId string) ([]domainmodels.Bet, error)
	GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.Bet, error)
}
