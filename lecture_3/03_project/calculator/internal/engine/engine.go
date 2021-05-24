package engine

import (
	"context"
	"log"
)

// Engine is the main component, responsible for consuming received bets and calculated bets,
// processing them and publishing the resulting bets.
type Engine struct {
	consumer  Consumer
	handler   Handler
	publisher Publisher
}

// New creates and returns a new engine.
func New(consumer Consumer, handler Handler, publisher Publisher) *Engine {
	return &Engine{
		consumer:  consumer,
		handler:   handler,
		publisher: publisher,
	}
}

// Start will run the engine.
func (e *Engine) Start(ctx context.Context) {
	err := e.processBets(ctx)
	if err != nil {
		log.Println("Engine failed to process bets received:", err)
		return
	}

	err = e.processEventsSettled(ctx)
	if err != nil {
		log.Println("Engine failed to process bets calculated:", err)
		return
	}

	<-ctx.Done()
}

func (e *Engine) processBets(ctx context.Context) error {
	consumedBets, err := e.consumer.ConsumeBets(ctx)
	if err != nil {
		return err
	}

	e.handler.HandleBets(ctx, consumedBets)

	return nil
}

func (e *Engine) processEventsSettled(ctx context.Context) error {
	consumedEventsSettled, err := e.consumer.ConsumeEventsSettled(ctx)
	if err != nil {
		return err
	}

	resultingBets := e.handler.HandleEventsSettled(ctx, consumedEventsSettled)
	e.publisher.PublishBetsCalculated(ctx, resultingBets)

	return nil
}
