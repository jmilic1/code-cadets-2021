package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http/offerfeeds"
)

func AxilisOfferFeed(httpClient *stdhttp.Client, updates chan models.Odd) *offerfeeds.AxilisOfferFeed {
	return offerfeeds.NewAxilisOfferFeed(httpClient, updates)
}
