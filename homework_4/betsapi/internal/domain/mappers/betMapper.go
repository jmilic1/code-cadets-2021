package mappers

import (
	dto "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers/models"
	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/infrastructure/sqlite/models"
)

// BetMapper maps storage bets to domain bets and vice versa.
type BetMapper struct{}

// NewBetMapper creates and returns a new BetMapper.
func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

// MapStorageBetToDomainBet maps the given storage bet into domain bet. Floating point values will
// be converted from corresponding integer values of the storage bet by dividing them with 100.
func (b *BetMapper) MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.Bet {
	return domainmodels.Bet{
		Id:                   storageBet.Id,
		CustomerId:           storageBet.CustomerId,
		Status:               storageBet.Status,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: float64(storageBet.SelectionCoefficient) / 100,
		Payment:              float64(storageBet.Payment) / 100,
		Payout:               float64(storageBet.Payout) / 100,
	}
}

// MapDomainBetToDto maps the given domain bet into dto bet.
func (b *BetMapper) MapDomainBetToDto(domainBet domainmodels.Bet) dto.BetResultDto {
	return dto.BetResultDto{
		Id:                   domainBet.Id,
		Status:               domainBet.Status,
		SelectionId:          domainBet.SelectionId,
		SelectionCoefficient: domainBet.SelectionCoefficient,
		Payment:              domainBet.Payment,
		Payout:               domainBet.Payout,
	}
}
