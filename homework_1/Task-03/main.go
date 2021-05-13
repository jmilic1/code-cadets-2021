package main

import (
	pokemonEncounterApi "code-cadets-2021/homework_1/Task-03/pokemon-encounter-api"
	"encoding/json"
	"flag"
	"github.com/pkg/errors"
	"log"
)

// main is the entrypoint for getting Pokemon information.
func main() {
	pokemon := flag.String("pokemon", "wooper", "Id or name of Pokemon to be retrieved")

	flag.Parse()

	if pokemon == nil {
		log.Fatal(
			"flag parsing returned nil pointer",
		)
	}

	encounters, err := pokemonEncounterApi.GetEncounters(*pokemon)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error while getting encounters"),
		)
	}

	output, err := json.Marshal(encounters)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error while marshaling pokemon encounters"),
		)
	}

	log.Println(string(output))
}
