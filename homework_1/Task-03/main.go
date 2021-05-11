package main

import (
	"encoding/json"
	"flag"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const url = "https://pokeapi.co/api/v2/pokemon/"

type pokemonDto struct {
	Name       string
	Encounters []string
}

type pokemon struct {
	Name                     string
	Location_area_encounters string
	Encounters               []string
}

type areaEncounters struct {
	Location_area struct {
		Name string
	}
}

func fetchFromUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyContent, nil
}

func main() {
	var id int
	var name string

	flag.IntVar(&id, "id", 0, "Id of Pokemon to be retrieved")
	flag.StringVar(&name, "name", "wooper", "Name of Pokemon to be retrieved")

	flag.Parse()

	var bodyContent []byte
	var err error

	if id == 0 {
		bodyContent, err = fetchFromUrl(url + name)
		if err != nil {
			log.Fatal(err, "Error while fetching content")
		}

	} else {
		bodyContent, err = fetchFromUrl(url + strconv.Itoa(id))
		if err != nil {
			log.Fatal(err, "Error while fetching content")
		}
	}

	var decodedPokemon pokemon
	err = json.Unmarshal(bodyContent, &decodedPokemon)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	bodyContent, err = fetchFromUrl(decodedPokemon.Location_area_encounters)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error while fetching content"),
		)
	}

	var decodedEncounters []areaEncounters
	err = json.Unmarshal(bodyContent, &decodedEncounters)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	dto := pokemonDto{
		Name: decodedPokemon.Name,
	}
	for _, val := range decodedEncounters {
		dto.Encounters = append(dto.Encounters, val.Location_area.Name)
	}

	output, err := json.Marshal(dto)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "marshalling the JSON body content"),
		)
	}

	log.Println(string(output))
}
