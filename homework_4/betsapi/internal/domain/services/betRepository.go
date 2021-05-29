package services

import (
	"context"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/infrastructure/sqlite/models"
)

type BetRepository interface {
	GetBetById(ctx context.Context, id string) (storagemodels.Bet, bool, error)
	GetBetsByCustomerId(ctx context.Context, customerId string) ([]storagemodels.Bet, error)
	GetBetsByStatus(ctx context.Context, status string) ([]storagemodels.Bet, error)
}
