package services

import (
	dto "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers/models"
	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/domain/models"
)

type BetMapper interface {
	MapDomainBetToDto(domainBet domainmodels.Bet) dto.BetResultDto
}
