# Pokedex CLI

A fun, interactive **command-line Pokédex** built in Go!  

Explore the Pokémon world, discover Pokémon in different areas, catch them, inspect their stats, and build your own Pokédex – all from your terminal.

This project uses the free [PokéAPI](https://pokeapi.co/) to fetch real Pokémon data.

## Features

- **Map navigation**: Use `map` to see the next 20 locations and `mapb` to go back
- **Explore areas**: `explore <area-name>` to list Pokémon that can appear there
- **Catch Pokémon**: `catch <pokemon-name>` – try your luck with a Poké Ball!
- **Inspect caught Pokémon**: `inspect <pokemon-name>` to view detailed stats
- **View your Pokédex**: `pokedex` to list all Pokémon you've caught
- **Help & Exit**: `help` for commands, `exit` to quit
- **Caching**: Responses from PokéAPI are cached in memory for faster repeated requests
- **Pagination**: Seamless next/previous navigation through the world map

## Tech Stack

- **Language**: Go (standard library + minimal dependencies)
- **API**: [PokéAPI](https://pokeapi.co/) (public, no key required)

## Prerequisites

- Go 1.22 or higher

## Installation & Running

1. Clone the repository:
   ```bash
   git clone https://github.com/CodeZeroSugar/pokedexcli.git
   cd pokedexcli
   ```

2. Build and run:
   ```bash
   go run .
   ```

   Or build a binary:
   ```bash
   go build -o pokedexcli
   ./pokedexcli
   ```

The program starts an interactive REPL prompt (`Pokedex >`).

## Commands

| Command              | Description                                      | Example                  |
|----------------------|--------------------------------------------------|--------------------------|
| `help`               | Show this list of commands                       | `help`                   |
| `exit`               | Exit the Pokédex                                 | `exit`                   |
| `map`                | Display the next page of location areas           | `map`                    |
| `mapb`               | Display the previous page of location areas       | `mapb`                   |
| `explore <area>`     | List Pokémon that can be encountered in an area  | `explore viridian-forest`|
| `catch <pokemon>`    | Attempt to catch a Pokémon                       | `catch pikachu`          |
| `inspect <pokemon>`  | View details of a caught Pokémon                 | `inspect pikachu`        |
| `pokedex`            | List all Pokémon in your Pokédex                 | `pokedex`                |

## Example Session

```
Pokedex > map
celadon-city
cerulean-cave
...
Pokedex > explore cerulean-cave
Possible encounters:
- golbat
- parasect
- raichu
- ...
Pokedex > catch raichu
Throwing a Pokeball at raichu...
raichu was caught!
Pokedex > inspect raichu
Name: raichu
Height: 8
Weight: 300
Stats:
  - hp: 60
  - attack: 90
  ...
Types:
  - electric
Pokedex > pokedex
You have caught these Pokemon:
- raichu
- pikachu
Pokedex > exit
```

## Notes

- Catching success depends on the Pokémon's base experience (easier for lower values)
- No persistence – your Pokédex is in-memory only (resets on exit)
- All API calls are cached to reduce load on PokéAPI

## Contributing

Feel free to open issues or pull requests! Ideas for improvements:
- Persistent storage (JSON file or SQLite)
- More commands (e.g., battles, evolution info)
- Colorful output with third-party libraries

## License

MIT License – use and modify freely!

Enjoy your Pokémon adventure in the terminal! 🐾
