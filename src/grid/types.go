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

func (cell GridCell) IsAlive() bool {
	return cell.Value == 1
}

func (cell GridCell) IsDead() bool {
	return !cell.IsAlive()
}

func (cell GridCell) Kill() {
	cell.Value = 0
}

func (cell GridCell) Resurrect() {
	cell.Value = 1
}

type delta struct {
	deltaRow int
	deltaCol int
}
