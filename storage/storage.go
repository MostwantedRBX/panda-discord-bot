package storage

import (
	"math/rand"
)

func ReturnRandomPokemon() string {
	//	TODO: find a better way to store this, seems like a mundane use of resources to have a package dedicated to this
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
