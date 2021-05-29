package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http/offerfeeds"
)

func HomeworkOfferFeed(httpClient *stdhttp.Client, updates chan models.Odd) *offerfeeds.HomeworkOfferFeed {
	return offerfeeds.NewHomeworkOfferFeed(httpClient, updates)
}
