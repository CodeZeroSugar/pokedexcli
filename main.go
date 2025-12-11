package main

type config struct {
	NextURL     string
	PreviousURL string
}

func main() {
	configURL := config{
		NextURL:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		PreviousURL: "",
	}

	startRepl(configURL)
}
