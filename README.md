# Pokedex
Interactive CLI game that simulates exploring the Pokemon world and catching Pokemon by making API calls to the [PokeApi](https://pokeapi.co/).

## Steps to run
1. Clone the repo
1. Run `go build`
1. Run `./pokedex`
1. Once the CLI is running, you can type `help` to show a series of commands that you can run. NOTE: `explore` and `inspect` command require an extra parameter, either a location or a pokemon.

A common flow for the CLI involves:
1. Run the `map` command to display some locations on the map
1. Choose a location and run `explore { location }` to list the Pokemon found in that location
1. Choose a Pokemon you want to catch and run `catch { pokemon }` to attempt a catch. 
1. If succesful, the Pokemon data will be stored in your Pokedex. Run `pokedex` list all Pokemon in your Pokedex and confirm your newly caught Pokemon has been stored. 
1. Choose a Pokemon from your Pokedex and run `inspect { pokemon }` to display basic stats about your Pokemon.
