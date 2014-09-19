package game

import (
	"EpiCraft/game/world"
	"time"
)

var (
	wordController *world.World
)

// GetWorld returns the only instance of the world.
func GetWorld() *world.World {
	return wordController
}

// Loop is the main loop of the game.
func Loop() {
	worldController := world.NewWorld()
	for {
		for _, entity := range worldController.Entities {
		}
		<-time.After(time.Millisecond * 33)
	}
}
