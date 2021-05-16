package bootstrap

import (
	"code-cadets-2021/homework_2/offerfeed/internal/domain/services"
)

func NewProcessingService(feed services.Feed, queue services.Queue) *services.FeedProcessorService {
	return services.NewFeedProcessorService(feed, queue)
}
