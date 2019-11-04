package grid

import "errors"

var deltaMap map[string]delta

func NewGrid(rows int, cols int) *Grid {
	deltaMap = map[string]delta{
		"tl": {-1, -1},
		"tc": {-1, 0},
		"tr": {-1, 1},
		"l":  {0, -1},
		"r":  {0, 1},
		"bl": {1, -1},
		"bc": {1, 0},
		"br": {1, 1},
	}

	grid := &Grid{
		grid: make([][]CellType, rows),

		dimensions: GridDimensions{
			rows: rows,
			cols: cols,
		},

		boundaries: GridBoundaries{
			top:    0,
			right:  cols - 1,
			bottom: rows - 1,
			left:   0,
		},
	}

	for index := range grid.grid {
		grid.grid[index] = make([]CellType, cols)
	}

	return grid
}

// Utility function for tests.
func (g Grid) initialize(table [][]CellType) {
	assert(len(table) == g.dimensions.rows, "Invalid initialize size (rows)")
	assert(len(table[0]) == g.dimensions.cols, "Invalid initialize size (cols)")

	for rowI, row := range table {
		for colI, value := range row {
			g.grid[rowI][colI] = value
		}
	}
}

func (g Grid) GetCell(pos GridCellPosition) CellType {
	assertRowBoundary(g, pos.Row)
	assertColBoundary(g, pos.Col)

	return g.grid[pos.Row][pos.Col]
}

func (g Grid) SetCell(value CellType, pos GridCellPosition) {
	assertRowBoundary(g, pos.Row)
	assertColBoundary(g, pos.Col)

	g.grid[pos.Row][pos.Col] = value // Should we store GridCells here?
}

func (g Grid) GetNeighbors(pos GridCellPosition) map[string]GridCell {
	neighbors := make(map[string]GridCell)

	neighbors["tl"] = g.getRelativeNeighbor(pos, "tl")
	neighbors["tc"] = g.getRelativeNeighbor(pos, "tc")
	neighbors["tr"] = g.getRelativeNeighbor(pos, "tr")
	neighbors["l"] = g.getRelativeNeighbor(pos, "l")
	neighbors["r"] = g.getRelativeNeighbor(pos, "r")
	neighbors["bl"] = g.getRelativeNeighbor(pos, "bl")
	neighbors["bc"] = g.getRelativeNeighbor(pos, "bc")
	neighbors["br"] = g.getRelativeNeighbor(pos, "br")

	return neighbors
}

func (g Grid) getRelativeNeighbor(pos GridCellPosition, deltaLabel string) GridCell {
	delta, exists := deltaMap[deltaLabel]
	if !exists {
		panic(errors.New("bad delta"))
	}

	neighbor := GridCell{
		pos: pos,
	}

	neighbor.pos.Row += delta.deltaRow
	neighbor.pos.Col += delta.deltaCol

	if neighbor.pos.Row < g.boundaries.top {
		neighbor.pos.Row = g.boundaries.bottom
	} else if neighbor.pos.Row > g.boundaries.bottom {
		neighbor.pos.Row = g.boundaries.top
	}

	if neighbor.pos.Col < g.boundaries.left {
		neighbor.pos.Col = g.boundaries.right
	} else if neighbor.pos.Col > g.boundaries.right {
		neighbor.pos.Col = g.boundaries.left
	}

	neighbor.value = g.GetCell(neighbor.pos)

	return neighbor
}

func (g Grid) IterateRows(cb func(index int, row []CellType)) {
	for index, row := range g.grid {
		cb(index, row)
	}
}

func (g Grid) IterateCells(cb func(row int, col int, value CellType)) {
	g.IterateRows(func (rowIndex int, row []CellType) {
		for colIndex, value := range row {
			cb(rowIndex, colIndex, value)
		}
	})
}

func (g Grid) Clone() *Grid {
	return NewGrid(g.dimensions.rows, g.dimensions.cols)
}