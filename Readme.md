# PokedexCLI 🐾

A simple, interactive **Pokedex command-line application** written in **Go**, inspired by the Pokémon universe.  
This project demonstrates clean Go project structure, API integration, caching, and a custom REPL.

---

## ✨ Features

- 📍 Explore Pokémon **location areas**
- 🗺️ List Pokémon available in a location
- 🔍 Inspect Pokémon details (stats, types, etc.)
- 🎒 Catch Pokémon and store them in your Pokédex
- 📖 View your captured Pokédex
- ⚡ In-memory **cache** to avoid redundant API calls
- 🧠 Interactive **REPL-style CLI**

---

## 🕹️ Available Commands

| Command              | Description                        |
| -------------------- | ---------------------------------- |
| `help`               | Show all available commands        |
| `map`                | List next set of Pokémon locations |
| `mapb`               | List previous set of locations     |
| `explore <location>` | List Pokémon in a location         |
| `catch <pokemon>`    | Attempt to catch a Pokémon         |
| `inspect <pokemon>`  | View Pokémon details               |
| `pokedex`            | View all caught Pokémon            |
| `exit`               | Exit the application               |

---

## 🧱 Project Structure

```

Pokedexcli/
├── internal/
│   ├── commands/               # CLI commands logic
│   │   ├── command_catch.go
│   │   ├── command_exit.go
│   │   ├── command_explore.go
│   │   ├── command_help.go
│   │   ├── command_inspect.go
│   │   ├── command_map.go
│   │   ├── command_pokedex.go
│   │   └── config.go
│   │
│   ├── pokeapi/                # PokeAPI request handlers
│   │   ├── pokeapi.go
│   │   ├── request_catch_pokemon.go
│   │   ├── request_explore_pokemon.go
│   │   └── request_location_area.go
│   │
│   ├── pokecache/              # In-memory cache
│   │   ├── pokecache.go
│   │   └── pokecache_test.go
│   │
│   └── shared/
│       └── types.go            # Shared data types
│
├── main.go                     # Program entry point
├── repl.go                     # REPL implementation
├── repl_test.go
├── go.mod
├── go.sum
└── README.md

````

---

## 🚀 Getting Started

### Prerequisites
- Go **1.20+**

### Clone the Repository
```bash
git clone https://github.com/kalim-Asim/Pokedexcli.git
cd Pokedexcli
````

### Run the Application

```bash
go run .
```

Or build a binary:

```bash
go build -o pokedexcli
./pokedexcli
```

---

## 🧪 Testing

Run all tests using:

```bash
go test ./...
```

---

## 🧠 Caching

The app uses a **custom in-memory cache** (`pokecache`) to:

* Reduce API calls
* Improve performance
* Demonstrate Go concurrency-safe design

---

## 📦 API Used

* **PokeAPI**: [https://pokeapi.co/](https://pokeapi.co/)

---

## 📌 Learning Goals

This project focuses on:

* Clean Go project layout
* HTTP API consumption
* CLI + REPL design
* State management
* Caching strategies
* Writing testable Go code

---

