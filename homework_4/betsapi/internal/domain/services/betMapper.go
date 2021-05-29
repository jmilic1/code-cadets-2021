package services

import (
	dto "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapStorageBetToDto(domainBet storagemodels.Bet) dto.BetResultDto
}
