package services

import (
	controllermodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapStorageBetToDto(domainBet storagemodels.Bet) controllermodels.BetResultDto
}
