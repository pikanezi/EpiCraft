package entities

import "math"

// Entity represents a player, a monster or an object (basically everything).
type Entity interface {
	SetPosX(int)
	SetPosY(int)
	SetAction(Action)

	GetPosX() int
	GetID() uint
	GetPosY() int
	GetType() Type
	GetStatus() Status
	GetRange(target Entity) float64

	Play() error
}

// SimpleEntity is the simplest entity.
type SimpleEntity struct {
	ID     uint   `json:"idGame,omitempty" bson:"idGame"`
	PosX   int    `json:"posX"`
	PosY   int    `json:"posY"`
	Type   Type   `json:"entityType,omitempty"`
	Status Status `json:"status,omitempty"`
	action Action
}

// GetID returns the identifier of the entity.
func (entity *SimpleEntity) GetID() uint {
	return entity.ID
}

// SetPosX sets the x position of the entity.
func (entity *SimpleEntity) SetPosX(x int) {
	entity.PosX = x
}

// GetPosX returns the x position of the entity.
func (entity *SimpleEntity) GetPosX() int {
	return entity.PosX
}

// SetPosY sets the y position of the entity.
func (entity *SimpleEntity) SetPosY(y int) {
	entity.PosY = y
}

// GetPosY returns the y position of the entity.
func (entity *SimpleEntity) GetPosY() int {
	return entity.PosY
}

// GetType returns the type of the Entity.
func (entity *SimpleEntity) GetType() Type {
	return entity.Type
}

// SetAction sets the action of the Entity.
func (entity *SimpleEntity) SetAction(action Action) {
	entity.action = action
}

// Play the action of the Entity and consumes it.
// If the Entity has no action it does nothing.
func (entity *SimpleEntity) Play() error {
	if entity.action == nil {
		return nil
	}
	err := entity.action.Play()
	entity.action = nil
	return err
}

// GetStatus return the statuses of the entity.
func (entity *SimpleEntity) GetStatus() Status {
	return entity.Status
}

// GetRange returns the range between the entity and the target in number of cells.
func (entity *SimpleEntity) GetRange(target Entity) float64 {
	x, y := float64(target.GetPosX()-entity.GetPosX()), float64(target.GetPosY()-entity.GetPosY())
	if x == 0 && y == 0 {
		return 0
	}
	return math.Sqrt(x*x + y*y)
}
