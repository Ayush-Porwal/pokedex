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

type PokemonWorldMapLocation struct {
    EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
    GameIndex            int64                 `json:"game_index"`
    ID                   int64                 `json:"id"`
    PokemonLocation      PokemonLocation       `json:"location"`
    Name                 string                `json:"name"`
    Names                []Name                `json:"names"`
    PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
    EncounterMethod PokemonLocation                    `json:"encounter_method"`
    VersionDetails  []EncounterMethodRateVersionDetail `json:"version_details"`
}

type EncounterMethodRateVersionDetail struct {
    Rate    int64           `json:"rate"`
    Version PokemonLocation `json:"version"`
}

type Name struct {
    Language PokemonLocation `json:"language"`
    Name     string          `json:"name"`
}

type PokemonEncounter struct {
    Pokemon        PokemonLocation                 `json:"pokemon"`
    VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

type PokemonEncounterVersionDetail struct {
    EncounterDetails []EncounterDetail `json:"encounter_details"`
    MaxChance        int64             `json:"max_chance"`
    Version          PokemonLocation   `json:"version"`
}

type EncounterDetail struct {
    Chance          int64           `json:"chance"`
    ConditionValues []interface{}   `json:"condition_values"`
    MaxLevel        int64           `json:"max_level"`
    Method          PokemonLocation `json:"method"`
    MinLevel        int64           `json:"min_level"`
}

type Pokemon struct {
    Abilities              []Ability     `json:"abilities"`
    BaseExperience         int64         `json:"base_experience"`
    Cries                  Cries         `json:"cries"`
    Forms                  []Species     `json:"forms"`
    GameIndices            []GameIndex   `json:"game_indices"`
    Height                 int64         `json:"height"`
    HeldItems              []HeldItem    `json:"held_items"`
    ID                     int64         `json:"id"`
    IsDefault              bool          `json:"is_default"`
    LocationAreaEncounters string        `json:"location_area_encounters"`
    Moves                  []Move        `json:"moves"`
    Name                   string        `json:"name"`
    Order                  int64         `json:"order"`
    PastAbilities          []interface{} `json:"past_abilities"`
    PastTypes              []interface{} `json:"past_types"`
    Species                Species       `json:"species"`
    Sprites                Sprites       `json:"sprites"`
    Stats                  []Stat        `json:"stats"`
    Types                  []Type        `json:"types"`
    Weight                 int64         `json:"weight"`
}

type Ability struct {
    Ability  Species `json:"ability"`
    IsHidden bool    `json:"is_hidden"`
    Slot     int64   `json:"slot"`
}

type Species struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type Cries struct {
    Latest string `json:"latest"`
    Legacy string `json:"legacy"`
}

type GameIndex struct {
    GameIndex int64   `json:"game_index"`
    Version   Species `json:"version"`
}

type HeldItem struct {
    Item           Species         `json:"item"`
    VersionDetails []VersionDetail `json:"version_details"`
}

type VersionDetail struct {
    Rarity  int64   `json:"rarity"`
    Version Species `json:"version"`
}

type Move struct {
    Move                Species              `json:"move"`
    VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
    LevelLearnedAt  int64   `json:"level_learned_at"`
    MoveLearnMethod Species `json:"move_learn_method"`
    VersionGroup    Species `json:"version_group"`
}

type GenerationV struct {
    BlackWhite Sprites `json:"black-white"`
}

type GenerationIv struct {
    DiamondPearl        Sprites `json:"diamond-pearl"`
    HeartgoldSoulsilver Sprites `json:"heartgold-soulsilver"`
    Platinum            Sprites `json:"platinum"`
}

type Versions struct {
    GenerationI    GenerationI     `json:"generation-i"`
    GenerationIi   GenerationIi    `json:"generation-ii"`
    GenerationIii  GenerationIii   `json:"generation-iii"`
    GenerationIv   GenerationIv    `json:"generation-iv"`
    GenerationV    GenerationV     `json:"generation-v"`
    GenerationVi   map[string]Home `json:"generation-vi"`
    GenerationVii  GenerationVii   `json:"generation-vii"`
    GenerationViii GenerationViii  `json:"generation-viii"`
}

type Other struct {
    DreamWorld      DreamWorld      `json:"dream_world"`
    Home            Home            `json:"home"`
    OfficialArtwork OfficialArtwork `json:"official-artwork"`
    Showdown        Sprites         `json:"showdown"`
}

type Sprites struct {
    BackDefault      string    `json:"back_default"`
    BackFemale       string    `json:"back_female"`
    BackShiny        string    `json:"back_shiny"`
    BackShinyFemale  *string   `json:"back_shiny_female"`
    FrontDefault     string    `json:"front_default"`
    FrontFemale      string    `json:"front_female"`
    FrontShiny       string    `json:"front_shiny"`
    FrontShinyFemale string    `json:"front_shiny_female"`
    Other            *Other    `json:"other,omitempty"`
    Versions         *Versions `json:"versions,omitempty"`
    Animated         *Sprites  `json:"animated,omitempty"`
}

type GenerationI struct {
    RedBlue RedBlue `json:"red-blue"`
    Yellow  RedBlue `json:"yellow"`
}

type RedBlue struct {
    BackDefault      string `json:"back_default"`
    BackGray         string `json:"back_gray"`
    BackTransparent  string `json:"back_transparent"`
    FrontDefault     string `json:"front_default"`
    FrontGray        string `json:"front_gray"`
    FrontTransparent string `json:"front_transparent"`
}

type GenerationIi struct {
    Crystal Crystal `json:"crystal"`
    Gold    Gold    `json:"gold"`
    Silver  Gold    `json:"silver"`
}

type Crystal struct {
    BackDefault           string `json:"back_default"`
    BackShiny             string `json:"back_shiny"`
    BackShinyTransparent  string `json:"back_shiny_transparent"`
    BackTransparent       string `json:"back_transparent"`
    FrontDefault          string `json:"front_default"`
    FrontShiny            string `json:"front_shiny"`
    FrontShinyTransparent string `json:"front_shiny_transparent"`
    FrontTransparent      string `json:"front_transparent"`
}

type Gold struct {
    BackDefault      string  `json:"back_default"`
    BackShiny        string  `json:"back_shiny"`
    FrontDefault     string  `json:"front_default"`
    FrontShiny       string  `json:"front_shiny"`
    FrontTransparent *string `json:"front_transparent,omitempty"`
}

type GenerationIii struct {
    Emerald          OfficialArtwork `json:"emerald"`
    FireredLeafgreen Gold            `json:"firered-leafgreen"`
    RubySapphire     Gold            `json:"ruby-sapphire"`
}

type OfficialArtwork struct {
    FrontDefault string `json:"front_default"`
    FrontShiny   string `json:"front_shiny"`
}

type Home struct {
    FrontDefault     string `json:"front_default"`
    FrontFemale      string `json:"front_female"`
    FrontShiny       string `json:"front_shiny"`
    FrontShinyFemale string `json:"front_shiny_female"`
}

type GenerationVii struct {
    Icons             DreamWorld `json:"icons"`
    UltraSunUltraMoon Home       `json:"ultra-sun-ultra-moon"`
}

type DreamWorld struct {
    FrontDefault string  `json:"front_default"`
    FrontFemale  *string `json:"front_female"`
}

type GenerationViii struct {
    Icons DreamWorld `json:"icons"`
}

type Stat struct {
    BaseStat int64   `json:"base_stat"`
    Effort   int64   `json:"effort"`
    Stat     Species `json:"stat"`
}

type Type struct {
    Slot int64   `json:"slot"`
    Type Species `json:"type"`
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

func ExploreArea(locationName string) (pokemonEncounters []PokemonEncounter, err error) {
    res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + locationName)

    if err != nil {
        return pokemonEncounters, err
    }

    body, err := io.ReadAll(res.Body)

    err = res.Body.Close()

    if err != nil {
        return pokemonEncounters, err
    }

    if res.StatusCode > 299 {
        log.Fatalf("Response failed with status code %d", res.StatusCode)
    }

    pokemonWorldMapLocation := PokemonWorldMapLocation{}

    err = json.Unmarshal(body, &pokemonWorldMapLocation)

    if err != nil {
        return pokemonEncounters, err
    }

    return pokemonWorldMapLocation.PokemonEncounters, err
}

func GetPokemon(pokemonName string) (pokemon Pokemon, err error) {
    res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemonName)

    if err != nil {
        return pokemon, err
    }

    body, err := io.ReadAll(res.Body)

    err = res.Body.Close()

    if err != nil {
        return pokemon, err
    }

    if res.StatusCode > 299 {
        log.Fatalf("Response failed with status code %d", res.StatusCode)
    }

    pokemon = Pokemon{}

    err = json.Unmarshal(body, &pokemon)

    if err != nil {
        return pokemon, err
    }

    return pokemon, err
}
