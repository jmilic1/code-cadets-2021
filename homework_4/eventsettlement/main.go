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
	response, err := client.Get(activeBetsPath)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error while sending GET request to Bets Api"),
		)
	}
	bodyContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error while reading response from Bets Api"),
		)
	}

	var decodedBets []betDto
	err = json.Unmarshal(bodyContent, &decodedBets)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error while decoding response from Bets Api"),
		)
	}

	selectionIds := map[string]bool{}
	for _, bet := range decodedBets {
		selectionIds[bet.SelectionId] = true
	}

	for selectionId := range selectionIds {
		var outcome string
		if rand.Float64() > 0.5 {
			outcome = "lost"
		} else {
			outcome = "won"
		}

		eventUpdate := &eventUpdateDto{
			Id:      selectionId,
			Outcome: outcome,
		}

		eventUpdateJson, err := json.Marshal(eventUpdate)
		if err != nil {
			log.Fatal(
				errors.WithMessage(err, "failed to marshal an event update"),
			)
		}

		reader := bytes.NewReader(eventUpdateJson)

		post, err := client.Post(eventPath, "text/json", reader)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(post)
	}
}
