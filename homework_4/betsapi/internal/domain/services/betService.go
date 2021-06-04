package services

import (
	"context"
	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers/models"
)

// BetService implements bet related functions.
type BetService struct {
	betRepository BetRepository
	betMapper     BetMapper
}

// NewBetService creates a new instance of BetService.
func NewBetService(betRepository BetRepository, betMapper BetMapper) *BetService {
	return &BetService{
		betRepository: betRepository,
		betMapper:     betMapper,
	}
}

// GetBet fetches and returns a bet based on given id.
func (b *BetService) GetBet(ctx context.Context, betId string) (models.BetResultDto, bool, error) {
	domainBet, found, err := b.betRepository.GetBetById(ctx, betId)
	if err != nil {
		return models.BetResultDto{}, false, err
	}

	dtoBet := b.betMapper.MapStorageBetToDto(domainBet)
	return dtoBet, found, nil
}

// GetBetsByCustomerId fetches and returns bets which are owned by given customerId
func (b *BetService) GetBetsByCustomerId(ctx context.Context, customerId string) ([]models.BetResultDto, error) {
	domainBets, err := b.betRepository.GetBetsByCustomerId(ctx, customerId)
	if err != nil {
		return []models.BetResultDto{}, err
	}

	var dtoBets []models.BetResultDto
	for _, bet := range domainBets {
		dtoBet := b.betMapper.MapStorageBetToDto(bet)
		dtoBets = append(dtoBets, dtoBet)
	}

	return dtoBets, nil
}

// GetBetsByStatus fetches and returns bets which have given status
func (b *BetService) GetBetsByStatus(ctx context.Context, status string) ([]models.BetResultDto, error) {
	domainBets, err := b.betRepository.GetBetsByStatus(ctx, status)
	if err != nil {
		return []models.BetResultDto{}, err
	}

	var dtoBets []models.BetResultDto
	for _, bet := range domainBets {
		dtoBet := b.betMapper.MapStorageBetToDto(bet)
		dtoBets = append(dtoBets, dtoBet)
	}

	return dtoBets, nil
}
