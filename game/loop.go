package game

import (
	"EpiCraft/game/world"
	"time"
)

var (
	worldController *world.World
)

// GetWorld returns the only instance of the world.
func GetWorld() *world.World {
	return worldController
}

// Loop is the main loop of the game.
func Loop() {
	worldController = world.NewWorld()
	for {
		for _, entity := range worldController.Players {
			entity.Play()
		}
		<-time.After(time.Millisecond * 33)
	}
}
