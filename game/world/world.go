package world

import "EpiCraft/game/entities"

// World holds everything.
type World struct {
	Zones   []*Zone
	Players []entities.Entity
}

// NewWorld returns a newly instantiated world.
func NewWorld() *World {
	return &World{
		Zones:   make([]*Zone, 0),
		Players: make([]entities.Entity, 0),
	}
}

// SetActionToPlayerByItsID finds the given entity by its id and sets an action.
func (w *World) SetActionToPlayerByItsID(id uint, action entities.Action) {
	for _, entity := range w.Players {
		if id == entity.GetID() {
			entity.SetAction(action)
		}
	}
}

// RemovePlayerByItsID finds the entity and removes it from the world.
func (w *World) RemovePlayerByItsID(id uint) bool {
	for i := 0; i < len(w.Players); i++ {
		if id == w.Players[i].GetID() {
			_, w.Players = w.Players[len(w.Players)-1], w.Players[:len(w.Players)-1]
			return true
		}
	}
	return false
}

// AddPlayer adds an entity to the entities of the World.
func (w *World) AddPlayer(entity entities.Entity) {
	w.Players = append(w.Players, entity)
}
