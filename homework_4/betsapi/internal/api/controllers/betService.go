package controllers

import (
	"github.com/gin-gonic/gin"

	dto "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers/models"
)

// BetService implements bet related functions.
type BetService interface {
	GetBet(ctx *gin.Context, betId string) (dto.BetResultDto, bool, error)
	GetBetsByCustomerId(ctx *gin.Context, customerId string) ([]dto.BetResultDto, error)
	GetBetsByStatus(ctx *gin.Context, betStatus string) ([]dto.BetResultDto, error)
}
