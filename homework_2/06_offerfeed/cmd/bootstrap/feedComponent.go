package bootstrap

import (
	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http"
)

func FeedComponent(updates chan models.Odd, feeds ...http.OfferFeed) (*http.FeedComponent, error) {
	return http.NewFeedComponent(updates, feeds)
}
