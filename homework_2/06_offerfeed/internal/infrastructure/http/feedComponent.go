package http

import (
	"context"
	"log"
	"sync"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
)

type FeedComponent struct {
	feeds   []OfferFeed
	updates chan models.Odd
}

type OfferFeed interface {
	// Start begins feeding offers to updates channel of this OfferFeed
	Start(ctx context.Context) error
	// GetUpdates returns updates channel of this OfferFeed
	GetUpdates() chan models.Odd
}

func NewFeedComponent(
	updates chan models.Odd,
	feeds []OfferFeed,
) (*FeedComponent, error) {
	return &FeedComponent{
		feeds:   feeds,
		updates: updates,
	}, nil
}

func (o *FeedComponent) Start(ctx context.Context) error {
	defer close(o.updates)

	feeds := o.feeds
	wg := sync.WaitGroup{}
	wg.Add(len(feeds))

	for _, feed := range feeds {
		go func(localFeed OfferFeed) {
			defer wg.Done()
			err := localFeed.Start(ctx)
			log.Printf(`"finished an offer feed with "%v" error`, err)
		}(feed)
	}

	wg.Wait()

	return nil
}

func (o *FeedComponent) GetUpdates() chan models.Odd {
	return o.updates
}
