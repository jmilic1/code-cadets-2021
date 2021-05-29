package offerfeeds

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	updates    chan models.Odd
	httpClient *http.Client
}

func NewAxilisOfferFeed(
	httpClient *http.Client,
	updates chan models.Odd,
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		httpClient: httpClient,
		updates:    updates,
	}
}

// Start reads Get http response from axilisFeedURL, serializes the data into an Odd model and sends it to updates channel
func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
			response, err := a.httpClient.Get(axilisFeedURL)
			if err != nil {
				return err
			}

			bodyContent, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}

			var decodedOfferOdds []axilisOfferOdd
			err = json.Unmarshal(bodyContent, &decodedOfferOdds)
			if err != nil {
				return err
			}

			for _, odd := range decodedOfferOdds {
				mappedOdd := models.Odd{
					Id:          odd.Id,
					Name:        odd.Name,
					Match:       odd.Match,
					Coefficient: odd.Details.Price,
					Timestamp:   time.Now(),
				}

				select {
				case <-ctx.Done():
					return nil
				case a.updates <- mappedOdd:
					// do nothing
				}
			}
		}
	}
}

func (a *AxilisOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}

// axilisOfferOdd models Get response from axilisFeedURL
type axilisOfferOdd struct {
	Id      string
	Name    string
	Match   string
	Details axilisOfferOddDetails
}

type axilisOfferOddDetails struct {
	Price float64
}
