package entities

// Type represents the type of entity.
type Type uint

const (
	// EntityStoneBlock represents a block of stone.
	EntityStoneBlock Type = iota + 1

	// EntityStoneWall represents a wall of stone.
	EntityStoneWall

	// EntityPlayerSpawner represents a spawn.
	EntityPlayerSpawner

	// EntityPlayer represents the player type.
	EntityPlayer

	// EntityMonster represents the monster type.
	EntityMonster
)
