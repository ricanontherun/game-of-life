package grid

type cellType int

type Grid struct {
	grid       [][]GridCell
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
	Value cellType
	Pos   GridCellPosition
}

type delta struct {
	deltaRow int
	deltaCol int
}
