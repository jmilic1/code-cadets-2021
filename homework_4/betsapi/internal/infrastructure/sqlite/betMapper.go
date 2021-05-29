package sqlite

import (
	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.Bet
}
