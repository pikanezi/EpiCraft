package entities

// Cell represents a single cell of a zone.
type Cell struct {
	PosX     int      `json:"posX"`
	PosY     int      `json:"posY"`
	Entities []Entity `json:"entities,omitempty"`
	Type     Type     `json:"type,omitempty"`
}

// NewCells returns a newly instantiated grid of cells.
func NewCells(xMax, yMax int) [][]*Cell {
	cells := make([][]*Cell, yMax)
	for i := 0; i < yMax; i++ {
		cells[i] = make([]*Cell, xMax)
	}
	return cells
}

// NewCell returns a new cell.
func NewCell(x, y int, cellType Type) *Cell {
	return &Cell{
		PosX:     x,
		PosY:     y,
		Entities: make([]Entity, 0),
		Type:     cellType,
	}
}
