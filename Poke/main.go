package main

import (
	"fmt"
	"math/rand"
	"time"

	p "./pokemon"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Hello world")
	
	 Factory := p.PokemonFactory{}
	 var myPokemons []p.Pokemon

	 for i := 0; i < 5; i++ {

		myPokemons = append(myPokemons, Factory.CreateRandomPokemon())
		 
	 }

	 for i := 0; i < len(myPokemons); i++ {
		 fmt.Printf("%v", myPokemons[i])
		 //println(myPokemons[i].Name)		 
	 }


}