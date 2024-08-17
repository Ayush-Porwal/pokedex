package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "github.com/Ayush-Porwal/pokedex/pokeapi"
    "log"
    "math/rand"
    "os"
    "strings"
    "time"
)

type pokedexCliCommand struct {
    name        string
    description string
}

type user struct {
    pokedex map[string]pokeapi.Pokemon
}

type config struct {
    next     *string
    previous *string
}

var pokedexCliCommands = map[string]pokedexCliCommand{
    "map": {
        name:        "map",
        description: "Explore the pokemon world! Prints 20 locations where the pokemon could be found.",
    },
    "mapb": {
        name:        "mapb",
        description: "Same as map command, except it shows the previous 20 locations shown before.",
    },
    "explore": {
        name:        "explore",
        description: "Explore the area, reveals the pokemons living in the area.",
    },
    "catch": {
        name:        "catch",
        description: "Catches the pokemon, the pokemon becomes available to the user's pokedex",
    },
    "inspect": {
        name:        "inspect",
        description: "It takes the name of a Pokemon as an argument. It should print the name, height, weight, stats and type(s) of the Pokemon",
    },
    "pokedex": {
        name:        "pokedex",
        description: "Lists all the pokemon which the user has caught",
    },
    "help": {
        name:        "help",
        description: "Displays a help message",
    },
    "exit": {
        name:        "exit",
        description: "Exit the pokedex",
    },
}

// max 5 since base_experience highest value is around 700
func difficultyLevel(experience int64) int64 {
    if experience < 100 {
        return 1
    } else if experience < 200 {
        return 2
    } else if experience < 300 {
        return 3
    } else if experience < 400 {
        return 4
    } else {
        return 5
    }
}

func main() {
    pokeapiURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
    scanner := bufio.NewScanner(os.Stdin)

    localConfig := config{
        next:     &pokeapiURL,
        previous: &pokeapiURL,
    }

    currentUser := user{
        pokedex: make(map[string]pokeapi.Pokemon),
    }

    cache := pokeapi.NewCache(7 * time.Second)

    for {
        fmt.Print("pokedex > ")

        for scanner.Scan() {
            scannedCommand := scanner.Text()

            if len(scannedCommand) != 0 {

                parts := strings.SplitN(scannedCommand, " ", 2)

                command := parts[0]

                if currentCliCommand, ok := pokedexCliCommands[command]; ok {

                    switch currentCliCommand.name {

                    case "map":
                        var err error
                        pokemonWorldMap := pokeapi.PokemonWorldMap{}

                        if response, found := cache.Get(*localConfig.next); found {
                            fmt.Println("Cache Hit")
                            cacheError := json.Unmarshal(response, &pokemonWorldMap)

                            if cacheError != nil {
                                log.Fatalf("Error!!: %s", cacheError)
                            }
                        } else {
                            fmt.Println("Cache Miss")

                            pokemonWorldMap, err = pokeapi.GetLocations(localConfig.next)

                            if err != nil {
                                log.Fatalf("Error!!: %s", err)
                            }
                            uncachedResponse, err := json.Marshal(pokemonWorldMap)

                            cache.Add(*localConfig.next, uncachedResponse)

                            if err != nil {
                                log.Fatalf("Error!!: %s", err)
                            }
                        }

                        for _, location := range pokemonWorldMap.Results {
                            fmt.Printf("%s\n", location.Name)
                        }

                        localConfig.next = pokemonWorldMap.Next

                        if pokemonWorldMap.Previous != nil {
                            localConfig.previous = pokemonWorldMap.Previous
                        } else {
                            localConfig.previous = &pokeapiURL
                        }

                        fmt.Println("")

                    case "mapb":
                        var err error
                        pokemonWorldMap := pokeapi.PokemonWorldMap{}

                        if response, found := cache.Get(*localConfig.previous); found {
                            fmt.Println("Cache Hit")
                            cacheError := json.Unmarshal(response, &pokemonWorldMap)

                            if cacheError != nil {
                                log.Fatalf("Error!!: %s", cacheError)
                            }
                        } else {
                            fmt.Println("Cache Miss")

                            pokemonWorldMap, err = pokeapi.GetLocations(localConfig.previous)

                            if err != nil {
                                log.Fatalf("Error!!: %s", err)
                            }

                            uncachedResponse, err := json.Marshal(pokemonWorldMap)

                            cache.Add(*localConfig.next, uncachedResponse)

                            if err != nil {
                                log.Fatalf("Error!!: %s", err)
                            }
                        }

                        for _, location := range pokemonWorldMap.Results {
                            fmt.Printf("%s\n", location.Name)
                        }

                        localConfig.next = pokemonWorldMap.Next

                        if pokemonWorldMap.Previous != nil {
                            localConfig.previous = pokemonWorldMap.Previous
                        } else {
                            localConfig.previous = &pokeapiURL
                        }

                        fmt.Println("")
                    case "explore":
                        var err error
                        argument := parts[1]
                        var pokemonEncounters []pokeapi.PokemonEncounter

                        if response, found := cache.Get(argument); found {
                            fmt.Println("Cache Hit")

                            cacheError := json.Unmarshal(response, &pokemonEncounters)

                            if cacheError != nil {
                                log.Fatalf("Error!!: %s", cacheError)
                            }
                        } else {
                            fmt.Println("Cache Miss")

                            pokemonEncounters, err = pokeapi.ExploreArea(argument)

                            if err != nil {
                                log.Fatalf("Error!!: %s", err)
                            }

                            uncachedResponse, err := json.Marshal(pokemonEncounters)

                            cache.Add(argument, uncachedResponse)

                            if err != nil {
                                log.Fatalf("Error!!: %s", err)
                            }
                        }

                        for _, pokemon := range pokemonEncounters {
                            fmt.Printf("%s\n", pokemon.Pokemon.Name)
                        }

                        fmt.Println("")

                    case "catch":
                        var err error
                        argument := parts[1]
                        var pokemon pokeapi.Pokemon

                        if _, ok := currentUser.pokedex[argument]; ok {
                            fmt.Println(fmt.Sprintf("You already own one %s", argument))
                            break
                        }

                        fmt.Println(fmt.Sprintf("Throwing a Pokeball at %s...", argument))

                        if response, found := cache.Get(argument); found {
                            cacheError := json.Unmarshal(response, &pokemon)

                            if cacheError != nil {
                                log.Fatalf("Error!!: %s", cacheError)
                            }
                        } else {
                            pokemon, err = pokeapi.GetPokemon(argument)

                            if err != nil {
                                log.Fatalf("Error!!: %s", err)
                            }

                            uncachedResponse, err := json.Marshal(pokemon)

                            cache.Add(argument, uncachedResponse)

                            if err != nil {
                                log.Fatalf("Error!!: %s", err)
                            }
                        }

                        probabilisticBaseExperience := rand.Int63n(pokemon.BaseExperience + 100*(5-difficultyLevel(pokemon.BaseExperience)))

                        if probabilisticBaseExperience > pokemon.BaseExperience {
                            currentUser.pokedex[argument] = pokemon
                            fmt.Println(fmt.Sprintf("%s was caught!!", argument))
                            fmt.Println("You may now inspect it with the inspect command.")
                        } else {
                            fmt.Println(fmt.Sprintf("%s escaped!", argument))
                        }
                    case "inspect":
                        argument := parts[1]
                        if pokemon, ok := currentUser.pokedex[argument]; ok {
                            fmt.Println(fmt.Sprintf("Name: %s", pokemon.Name))
                            fmt.Println(fmt.Sprintf("Height: %d", pokemon.Height))
                            fmt.Println(fmt.Sprintf("Weight: %d", pokemon.Weight))
                            fmt.Println("Stats:")
                            for _, stat := range pokemon.Stats {
                                fmt.Println(fmt.Sprintf(" -%s:%d", stat.Stat.Name, stat.BaseStat))
                            }
                            fmt.Println("Types:")
                            for _, pokemonType := range pokemon.Types {
                                fmt.Println(fmt.Sprintf(" -%s", pokemonType.Type.Name))
                            }
                        } else {
                            fmt.Println("You haven't caught this pokemon yet...")
                            fmt.Println("")
                        }
                        fmt.Println("")
                    case "pokedex":
                        fmt.Println("Your Pokedex: ")
                        for pokemon, _ := range currentUser.pokedex {
                            fmt.Println(fmt.Sprintf("- %s", pokemon))
                        }
                        fmt.Println("")
                    case "help":
                        fmt.Println("\nWelcome to the Pokedex!\nUsage:")
                        fmt.Println("")
                        for _, cliCommand := range pokedexCliCommands {
                            fmt.Println(cliCommand.name, " : ", cliCommand.description)
                        }
                        fmt.Println("")

                    case "exit":
                        os.Exit(0)

                    default:
                        fmt.Println("Warning!! No such command")
                    }

                    break
                }
            }
        }
    }
}
