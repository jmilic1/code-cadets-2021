package mappers

import (
	dto "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/infrastructure/sqlite/models"
)

// BetMapper maps storage bets to dto bets and vice versa.
type BetMapper struct{}

// NewBetMapper creates and returns a new BetMapper.
func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

// MapStorageBetToDto maps the given storage bet into a dto bet.
func (b *BetMapper) MapStorageBetToDto(storageBet storagemodels.Bet) dto.BetResultDto {
	return dto.BetResultDto{
		Id:                   storageBet.Id,
		Status:               storageBet.Status,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: float64(storageBet.SelectionCoefficient) / 100,
		Payment:              float64(storageBet.Payment) / 100,
		Payout:               float64(storageBet.Payout) / 100,
	}
}
