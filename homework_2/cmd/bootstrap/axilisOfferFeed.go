package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
)

func NewAxilisOfferFeed() *http.AxilisOfferFeed {
	return http.NewAxilisOfferFeed(stdhttp.Client{})
}
