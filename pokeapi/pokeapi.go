package pokeapi

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
)

type PokemonLocation struct {
    Name string `json:"name"`
    Url  string `json:"url"`
}

type PokemonWorldMap struct {
    Count    int64             `json:"count"`
    Next     *string           `json:"next"`
    Previous *string           `json:"previous"`
    Results  []PokemonLocation `json:"results"`
}

func GetLocations(url *string) (pokemonWorldMap PokemonWorldMap, err error) {

    res, err := http.Get(*url)

    if err != nil {
        return pokemonWorldMap, err
    }

    body, err := io.ReadAll(res.Body)

    err = res.Body.Close()

    if err != nil {
        return pokemonWorldMap, err
    }

    if res.StatusCode > 299 {
        log.Fatalf("Response failed with status code %d", res.StatusCode)
    }

    err = json.Unmarshal(body, &pokemonWorldMap)

    if err != nil {
        return pokemonWorldMap, err
    }

    return pokemonWorldMap, err
}
