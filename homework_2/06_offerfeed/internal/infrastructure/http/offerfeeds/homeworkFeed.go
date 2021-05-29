package offerfeeds

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
)

const axilisFeedHomeworkURL = "http://18.193.121.232/axilis-feed-2"

type HomeworkOfferFeed struct {
	updates    chan models.Odd
	httpClient *http.Client
}

func NewHomeworkOfferFeed(
	httpClient *http.Client,
	updates chan models.Odd,
) *HomeworkOfferFeed {
	return &HomeworkOfferFeed{
		httpClient: httpClient,
		updates:    updates,
	}
}

// Start reads Get http response from axilisFeedHomeworkURL and sends it to updates channel
func (h *HomeworkOfferFeed) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
			response, err := h.httpClient.Get(axilisFeedHomeworkURL)
			if err != nil {
				return err
			}

			bodyContent, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}

			rows := strings.Split(string(bodyContent), "\n")
			for _, row := range rows {
				odd, err := parseOdd(row)
				if err != nil {
					continue
				}

				select {
				case <-ctx.Done():
					return nil
				case h.updates <- odd:
					// do nothing
				}
			}
		}
	}
}

func (h *HomeworkOfferFeed) GetUpdates() chan models.Odd {
	return h.updates
}

func parseOdd(row string) (models.Odd, error) {
	fields := strings.Split(row, ",")

	id := fields[0][2:]
	name := fields[1]
	match := fields[2]
	timestamp := time.Now()
	coefficient, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return models.Odd{}, err
	}

	odd := models.Odd{
		Id:          id,
		Name:        name,
		Match:       match,
		Coefficient: coefficient,
		Timestamp:   timestamp,
	}
	return odd, nil
}
