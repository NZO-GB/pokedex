# Pokédex

> An interactive command-line Pokédex built with Go and the public
> PokéAPI.

Pokédex is a terminal application that lets you explore the Pokémon
world from your shell. Browse location areas, discover the Pokémon that
inhabit them, attempt to catch them, inspect their stats, and build your
own in-memory Pokédex.

The project showcases API consumption, HTTP clients, JSON decoding,
caching, REPL design, and command-based application architecture in Go.

------------------------------------------------------------------------

## Features

-   Interactive REPL interface
-   Browse Pokémon location areas with pagination
-   Explore individual areas to discover wild Pokémon
-   Catch Pokémon with a probability-based capture mechanic
-   Inspect detailed Pokémon statistics
-   Maintain your own Pokédex during a session
-   In-memory caching to reduce duplicate API requests
-   Automatic HTTP timeout handling

------------------------------------------------------------------------

## Tech Stack

  Technology     Purpose
  -------------- --------------------------
  Go             CLI application
  PokéAPI        Pokémon data source
  HTTP           API communication
  JSON           Response parsing
  Custom cache   Avoid repeated API calls

------------------------------------------------------------------------

## Installation

### Prerequisites

-   Go 1.22+

### Clone

``` bash
git clone https://github.com/NZO-GB/pokedex.git
cd pokedex
```

### Run

``` bash
go run .
```

or build an executable:

``` bash
go build -o pokedex
./pokedex
```

------------------------------------------------------------------------

## Commands

  Command               Description
  --------------------- ---------------------------------------------
  `help`                Display available commands
  `map`                 Show the next page of location areas
  `mapb`                Show the previous page
  `explore <area>`      List Pokémon found in a location area
  `catch <pokemon>`     Attempt to catch a Pokémon
  `inspect <pokemon>`   Display detailed stats for a caught Pokémon
  `pokedex`             List every Pokémon you've caught
  `exit`                Exit the application

------------------------------------------------------------------------

## Example Session

``` text
> map
> explore canalave-city-area
> catch pikachu
> inspect pikachu
> pokedex
```

------------------------------------------------------------------------

## Architecture

``` text
User
 │
 ▼
REPL
 │
 ▼
Command Dispatcher
 │
 ├────────► HTTP Client
 │            │
 │            ▼
 │        PokéAPI
 │
 └────────► Cache
```

------------------------------------------------------------------------

## Project Structure

``` text
.
├── internal/
│   ├── cache/
│   └── pokeapi/
├── command_*.go
├── repl.go
├── main.go
└── go.mod
```

------------------------------------------------------------------------

## Learning Goals

This project demonstrates:

-   Designing a REPL in Go
-   Working with REST APIs
-   JSON unmarshalling
-   HTTP client usage
-   Package organization
-   Simple caching strategies
-   State management across commands

------------------------------------------------------------------------

## Future Improvements

-   Persist the Pokédex between sessions
-   Colored terminal output
-   Search and filtering
-   Better battle/capture mechanics
-   Unit and integration tests
-   Fuzzy command matching
