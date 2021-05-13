package encounters

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// pokeApiUrl contains base path to pokeAPI
const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

// pokemonData is a basic model of Pokemon information
// defined by json response for "https://pokeapi.co/api/v2/pokemon/1"
type pokemonData struct {
	Name          string `json:"name"`
	EncountersUrl string `json:"location_area_encounters"`
}

// PokemonEncounters contains a pokemon's name and where it can be encountered
type PokemonEncounters struct {
	Name       string
	Encounters []string
}

// encounterData is a basic model containing information about where a Pokemon can be found
// defined by json response for "https://pokeapi.co/api/v2/pokemon/1/encounters"
type encounterData struct {
	LocationArea struct {
		Name string `json:"name"`
	} `json:"location_area"`
}

// fetchFromUrl fetches data from given url
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

// fetchEncounters fetches encounterData from given url
func fetchEncounters(url string) ([]encounterData, error) {
	bodyContent, err := fetchFromUrl(url)
	if err != nil {
		return nil, err
	}

	var decodedEncounters []encounterData
	err = json.Unmarshal(bodyContent, &decodedEncounters)
	if err != nil {
		return nil, err
	}

	return decodedEncounters, nil
}

// fetchPokemon fetches pokemonData data from given url
func fetchPokemon(url string) (pokemonData, error) {
	bodyContent, err := fetchFromUrl(url)
	if err != nil {
		return pokemonData{}, err
	}

	var decodedPokemon pokemonData
	err = json.Unmarshal(bodyContent, &decodedPokemon)
	if err != nil {
		return pokemonData{}, err
	}

	return decodedPokemon, nil
}

// GetEncounters returns encounters for given pokemon
func GetEncounters(pokemon string) (PokemonEncounters, error) {
	u, err := url.Parse(pokeApiUrl)
	if err != nil {
		return PokemonEncounters{}, err
	}

	u.Path = path.Join(u.Path, pokemon)
	completeUrl := u.String()

	fetchedPokemon, err := fetchPokemon(completeUrl)
	if err != nil {
		return PokemonEncounters{}, err
	}

	completeUrl = fetchedPokemon.EncountersUrl

	fetchedEncounters, err := fetchEncounters(completeUrl)
	if err != nil {
		return PokemonEncounters{}, err
	}

	output := PokemonEncounters{
		Name: fetchedPokemon.Name,
	}
	for _, encounter := range fetchedEncounters {
		output.Encounters = append(output.Encounters, encounter.LocationArea.Name)
	}

	return output, nil
}
