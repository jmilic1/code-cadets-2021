package handler

import (
	"context"
	"log"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Handler handles bets received and bets calculated.
type Handler struct {
	betCalculatedRepository BetCalculatedRepository
}

// New creates and returns a new Handler.
func New(betRepository BetCalculatedRepository) *Handler {
	return &Handler{
		betCalculatedRepository: betRepository,
	}
}

// HandleBets handles bets.
func (h *Handler) HandleBets(
	ctx context.Context,
	bets <-chan rabbitmqmodels.Bet,
) {
	go func() {
		for bet := range bets {
			// Calculate the domain bet based on the incoming bet.
			domainBet := domainmodels.BetCalculated{
				Id:                   bet.Id,
				SelectionId:          bet.SelectionId,
				SelectionCoefficient: bet.SelectionCoefficient,
				Payment:              bet.Payment,
			}

			_, found, err := h.betCalculatedRepository.GetBetCalculatedByID(ctx, bet.Id)
			if err != nil {
				log.Println("Failed to query bet, error: ", err)
				continue
			}
			// Ignore logging if bet already exists.
			if found {
				continue
			}

			log.Println("Processing bet, betId:", bet.Id)
			// Insert the domain bet into the repository.
			err = h.betCalculatedRepository.InsertBetCalculated(ctx, domainBet)
			if err != nil {
				log.Println("Failed to insert bet, error: ", err)
				continue
			}
		}
	}()
}

// HandleEventsSettled handles event updates.
func (h *Handler) HandleEventsSettled(
	ctx context.Context,
	eventUpdates <-chan rabbitmqmodels.EventSettled,
) <-chan rabbitmqmodels.BetCalculated {
	resultingBets := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(resultingBets)

		for eventUpdate := range eventUpdates {
			log.Println("Processing settled event, eventId:", eventUpdate.Id)

			// Fetch the domain bet.
			domainBets, exists, err := h.betCalculatedRepository.GetBetBySelectionID(ctx, eventUpdate.Id)
			if err != nil {
				log.Println("Failed to fetch bets which should be updated, error: ", err)
				continue
			}
			if !exists {
				log.Println("Bets with selectionId do not exist, selectionId: ", eventUpdate.Id)
				continue
			}

			for _, bet := range domainBets {
				var payout float64
				if eventUpdate.Outcome == "won" {
					payout = bet.Payment * bet.SelectionCoefficient
				} else {
					payout = 0
				}

				// Calculate the resulting bet, which should be published.
				resultingBet := rabbitmqmodels.BetCalculated{
					Id:     bet.Id,
					Status: eventUpdate.Outcome,
					Payout: payout,
				}

				log.Printf("Returning bet with id: %s, outcome: %s, payout: %f", resultingBet.Id, eventUpdate.Outcome, resultingBet.Payout)

				select {
				case resultingBets <- resultingBet:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return resultingBets
}
