package entities

var (
	ids          uint
	availableIDs = make([]uint, 0)
)

// NewEntity returns a new simple entity.
func NewEntity(posX, posY int, entityType Type) Entity {
	ids++
	switch entityType {
	case EntityPlayerSpawner, EntityStoneBlock, EntityStoneWall:
		return &SimpleEntity{
			PosX: posX,
			PosY: posY,
			Type: entityType,
			ID:   getAvailableID(),
		}
	}
	panic("unknown type")
}

// DestroyEntity free its id.
func DestroyEntity(entity Entity) {
	availableIDs = append(availableIDs, entity.GetID())
}

// getAvailableID returns the first available id.
func getAvailableID() (id uint) {
	if len(availableIDs) != 0 {
		id, availableIDs = availableIDs[len(availableIDs)-1], availableIDs[:len(availableIDs)-1]
	} else {
		id = ids
		ids++
	}
	return id
}
