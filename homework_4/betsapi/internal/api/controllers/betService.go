package controllers

import (
	"context"
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers/models"
)

// BetService implements bet related functions.
type BetService interface {
	GetBet(ctx context.Context, betId string) (models.BetResultDto, bool, error)
	GetBetsByCustomerId(ctx context.Context, customerId string) ([]models.BetResultDto, error)
	GetBetsByStatus(ctx context.Context, betStatus string) ([]models.BetResultDto, error)
}
