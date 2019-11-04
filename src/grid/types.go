package grid

type CellType int

type Grid struct {
	grid       [][]CellType
	dimensions GridDimensions
	boundaries GridBoundaries
}

type GridDimensions struct {
	rows int
	cols int
}

type box struct {
	top    int
	right  int
	bottom int
	left   int
}

type GridBoundaries box

type GridCellPosition struct {
	Row int
	Col int
}

type GridCell struct {
	value CellType
	pos   GridCellPosition
}

type delta struct {
	deltaRow int
	deltaCol int
}