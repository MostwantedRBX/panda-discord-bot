package storage

import (
	"math/rand"
)

// Random stuffs

func ReturnRandomPokemon() string {
	pokemon := [6]string{
		"Whose that pokemon!",
		"Its Pikachu!",
		"I choose you, Team Rocket!",
		"I'm pretty sure Ash has been 10 years old for 20 years now...",
		"My favourite Pokemon is Snorlax",
		"I heard there was a panda Pokemon...",
	}
	return pokemon[rand.Intn(5)]
}
