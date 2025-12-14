package pokeapi

type LocationResults struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon
}

type Pokemon struct {
	Name string
	URL  string
}

type PokeInfo struct {
	BaseExperience int `json:"base_experience"`
	Name           string
	Height         int
	Weight         int
	Stats          []PokeStats
	Types          []PokeTypes
}

type PokeStats struct {
	Value int      `json:"base_stat"`
	Stat  StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type PokeTypes struct {
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}
