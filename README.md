# PokedexCLI

A command-line interface (CLI) application for exploring the world of Pokémon. This tool allows you to discover new areas, encounter different Pokémon, and build your own personal Pokédex.

## Features

*   **Explore Pokémon World**: Discover location areas and the Pokémon that inhabit them.
*   **Catch Pokémon**: Try your luck at catching Pokémon to add to your collection.
*   **Personal Pokédex**: View the Pokémon you've successfully caught.
*   **Inspect Pokémon**: Get detailed information about the Pokémon in your Pokédex, including their stats and types.
*   **Built-in Caching**: Caches API responses to provide a faster experience and reduce API calls.

## Commands

The following commands are available in the PokedexCLI:

*   `help`: Displays a help message with all available commands.
*   `map`: Lists the next set of location areas.
*   `mapb`: Lists the previous set of location areas.
*   `explore <area_name>`: Lists the Pokémon found in a specific area.
*   `catch <pokemon_name>`: Attempts to catch a Pokémon.
*   `inspect <pokemon_name>`: Displays information about a caught Pokémon.
*   `pokedex`: Lists all the Pokémon you have caught.
*   `exit`: Exits the PokedexCLI.

## Getting Started

### Prerequisites

*   Go (Golang) installed on your system.

### Building from Source

1.  Clone the repository:
    ```bash
    git clone https://github.com/lucasrodlima/pokedexcli.git
    cd pokedexcli
    ```

2.  Build the application:
    ```bash
    go build
    ```

3.  Run the application:
    ```bash
    ./pokedexcli
    ```

## Usage

Once the application is running, you can start using the commands listed above. For example, to see the available location areas, use the `map` command:

```
Pokedex > map
```

To explore an area and see which Pokémon are there, use the `explore` command followed by the area name:

```
Pokedex > explore canalave-city-area
```
