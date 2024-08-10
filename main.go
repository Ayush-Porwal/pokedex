package main

import (
    "bufio"
    "fmt"
    "github.com/Ayush-Porwal/pokedex/pokeapi"
    "log"
    "os"
)

type pokedexCliCommand struct {
    name        string
    description string
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
    "help": {
        name:        "help",
        description: "Displays a help message",
    },
    "exit": {
        name:        "exit",
        description: "Exit the pokedex",
    },
}

func main() {
    pokeapiURL := "https://pokeapi.co/api/v2/location-area"
    scanner := bufio.NewScanner(os.Stdin)
    localConfig := config{
        next:     &pokeapiURL,
        previous: &pokeapiURL,
    }

    for {
        fmt.Print("pokedex > ")

        for scanner.Scan() {
            scannedCommand := scanner.Text()

            if len(scannedCommand) != 0 {

                if currentCliCommand, ok := pokedexCliCommands[scannedCommand]; ok {

                    switch currentCliCommand.name {
                    case "map":
                        pokemonWorldMap, err := pokeapi.GetLocations(localConfig.next)

                        if err != nil {
                            log.Fatalf("Error!!: %s", err)
                        }

                        for _, location := range pokemonWorldMap.Results {
                            fmt.Printf("%s\n", location.Name)
                        }

                        localConfig.next = pokemonWorldMap.Next
                        localConfig.previous = pokemonWorldMap.Previous
                        fmt.Println("")

                    case "mapb":
                        pokemonWorldMap, err := pokeapi.GetLocations(localConfig.previous)

                        if err != nil {
                            log.Fatalf("Error!!: %s", err)
                        }

                        for _, location := range pokemonWorldMap.Results {
                            fmt.Printf("%s\n", location.Name)
                        }

                        localConfig.next = pokemonWorldMap.Next
                        localConfig.previous = pokemonWorldMap.Previous
                        fmt.Println("")

                    case "help":
                        fmt.Println("\nWelcome to the Pokedex!\nUsage:\n")
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
