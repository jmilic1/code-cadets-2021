package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const activeBetsPath = "http://127.0.0.1:8082/bets?status=active"
const eventPath = "http://127.0.0.1:8081/event/update"
const lostOutcome = "lost"
const wonOutcome = "won"

type betDto struct {
	Id                   string  `json:"id"`
	Status               string  `json:"status"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
	Payout               float64 `json:"payout"`
}

type eventUpdateDto struct {
	Id      string `json:"id"`
	Outcome string `json:"outcome"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	client := http.Client{Timeout: time.Second * 5}
	decodedBets, err := getActiveBets(client)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error while fetching active bets"),
		)
	}

	selectionIds := map[string]bool{}
	for _, bet := range decodedBets {
		selectionIds[bet.SelectionId] = true
	}

	for selectionId := range selectionIds {
		var outcome string
		if rand.Float64() > 0.5 {
			outcome = lostOutcome
		} else {
			outcome = wonOutcome
		}

		eventUpdate := &eventUpdateDto{
			Id:      selectionId,
			Outcome: outcome,
		}

		err = sendEventUpdate(eventUpdate, client)
		if err != nil {
			log.Fatal(
				errors.WithMessage(err, "error while sending event updates"),
			)
		}
	}
}

// getActiveBets retrieves bets which have active status from Bets Api.
func getActiveBets(client http.Client) ([]betDto, error) {
	response, err := client.Get(activeBetsPath)
	if err != nil {
		return nil, errors.WithMessage(err, "error while sending GET request to Bets Api")
	}

	bodyContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "error while reading response from Bets Api")
	}

	var decodedBets []betDto
	err = json.Unmarshal(bodyContent, &decodedBets)
	if err != nil {
		return nil, errors.WithMessage(err, "error while decoding response from Bets Api")
	}

	return decodedBets, nil
}

func sendEventUpdate(eventUpdate *eventUpdateDto, client http.Client) error {
	eventUpdateJson, err := json.Marshal(eventUpdate)
	if err != nil {
		return errors.WithMessage(err, "failed to marshal an event update")
	}

	reader := bytes.NewReader(eventUpdateJson)

	_, err = client.Post(eventPath, "text/json", reader)
	if err != nil {
		return errors.WithMessage(err, "error while sending POST request to event Api")
	}
	return nil
}
