package pokemon

import (
	"math/rand"
)


type Pokemon struct {
	Name        string
	PokemonType string
}

type PokemonFactory struct {
}

func (p *PokemonFactory) CreateRandomPokemon() Pokemon {

	
	names := []string{"Pikachu", "Bulbasaur", "Charmander", "Evee"}
	types := map[string]string{
		"Pikachu":    "Lighting",
		"Bulbasaur":  "Grass",
		"Charmander": "Fire",
		"Evee":       "Normal",
	}

	poke := new(Pokemon)
	poke.Name = names[rand.Intn(len(names))]
	poke.PokemonType = types[poke.Name]

	return *poke
}