package controllers

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/api/controllers/models"
)

// BetService implements bet related functions.
type BetService interface {
	SendBet(betRequest models.BetRequestDto) error
}
