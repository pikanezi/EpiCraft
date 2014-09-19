package world

import "EpiCraft/game/entities"

const (
	// Zone0 is the name of the first zone.
	Zone0 = "zone_0"
)

// Zone represents an entire zone of the world.
type Zone struct {
	Name         string             `json:"name,omitempty"`
	XMax         int                `json:"xMax,omitempty"`
	YMax         int                `json:"yMax,omitempty"`
	PlayerNumber int                `json:"playerNumber,omitempty"`
	Cells        [][]*entities.Cell `json:"cells,omitempty"`
	Entities     []*entities.Entity `json:"-"`
	spawn        *entities.Cell
}
