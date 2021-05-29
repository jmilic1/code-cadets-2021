package services

import (
	"github.com/gin-gonic/gin"

	dto "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/api/controllers/models"
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
func (b *BetService) GetBet(ctx *gin.Context, betId string) (dto.BetResultDto, bool, error) {
	domainBet, found, err := b.betRepository.GetBetById(ctx, betId)
	if err != nil {
		return dto.BetResultDto{}, false, err
	}

	dtoBet := b.betMapper.MapDomainBetToDto(domainBet)
	return dtoBet, found, nil
}

// GetBetsByCustomerId fetches and returns bets which are owned by given customerId
func (b *BetService) GetBetsByCustomerId(ctx *gin.Context, customerId string) ([]dto.BetResultDto, error) {
	domainBets, err := b.betRepository.GetBetsByCustomerId(ctx, customerId)
	if err != nil {
		return []dto.BetResultDto{}, err
	}

	var dtoBets []dto.BetResultDto
	for _, bet := range domainBets {
		dtoBet := b.betMapper.MapDomainBetToDto(bet)
		dtoBets = append(dtoBets, dtoBet)
	}

	return dtoBets, nil
}

// GetBetsByStatus fetches and returns bets which have given status
func (b *BetService) GetBetsByStatus(ctx *gin.Context, status string) ([]dto.BetResultDto, error) {
	domainBets, err := b.betRepository.GetBetsByStatus(ctx, status)
	if err != nil {
		return []dto.BetResultDto{}, err
	}

	var dtoBets []dto.BetResultDto
	for _, bet := range domainBets {
		dtoBet := b.betMapper.MapDomainBetToDto(bet)
		dtoBets = append(dtoBets, dtoBet)
	}

	return dtoBets, nil
}
