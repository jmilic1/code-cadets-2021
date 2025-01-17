package services

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"context"
)

type FeedProcessorService struct {
	feed  Feed
	queue Queue
}

func NewFeedProcessorService(argFeed Feed, argqueue Queue) *FeedProcessorService {
	return &FeedProcessorService{
		feed:  argFeed,
		queue: argqueue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	feedChannel := f.feed.GetUpdates()
	queueChannel := f.queue.GetSource()
	defer close(queueChannel)

	for val := range feedChannel {
		val.Coefficient *= 2
		queueChannel <- val
	}

	return nil
}

type Feed interface {
	GetUpdates() chan models.Odd
}

type Queue interface {
	GetSource() chan models.Odd
}
