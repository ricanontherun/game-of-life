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
		grid: make([][]GridCell, rows),

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
		grid.grid[index] = make([]GridCell, cols)
	}

	return grid
}

// Utility function for tests.
func (g Grid) initialize(table [][]cellType) {
	assert(len(table) == g.dimensions.rows, "Invalid initialize size (rows)")
	assert(len(table[0]) == g.dimensions.cols, "Invalid initialize size (cols)")

	for rowI, row := range table {
		for colI, _ := range row {
			g.SetCell(GridCell{
				Value: table[rowI][colI],
				Pos:   GridCellPosition{rowI, colI},
			})
		}
	}
}

func (g Grid) GetCell(pos GridCellPosition) GridCell {
	assertRowBoundary(g, pos.Row)
	assertColBoundary(g, pos.Col)

	return g.grid[pos.Row][pos.Col]
}

func (g Grid) SetCell(cell GridCell) {
	assertRowBoundary(g, cell.Pos.Row)
	assertColBoundary(g, cell.Pos.Col)

	g.grid[cell.Pos.Row][cell.Pos.Col] = cell
}

func (g Grid) GetNeighbors(cell GridCell) map[string]GridCell {
	neighbors := make(map[string]GridCell)

	pos := cell.Pos

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

	neighborPosition := pos

	// Adjust the neighbors position on the grid.
	neighborPosition.Row += delta.deltaRow
	neighborPosition.Col += delta.deltaCol

	// Wrap any values which have left the board.
	if neighborPosition.Row < g.boundaries.top {
		neighborPosition.Row = g.boundaries.bottom
	} else if neighborPosition.Row > g.boundaries.bottom {
		neighborPosition.Row = g.boundaries.top
	}

	if neighborPosition.Col < g.boundaries.left {
		neighborPosition.Col = g.boundaries.right
	} else if neighborPosition.Col > g.boundaries.right {
		neighborPosition.Col = g.boundaries.left
	}

	return g.GetCell(neighborPosition)
}

func (g Grid) IterateRows(cb func(index int, row []GridCell)) {
	for index, row := range g.grid {
		cb(index, row)
	}
}

func (g Grid) IterateCells(cb func(cell GridCell)) {
	g.IterateRows(func(rowIndex int, row []GridCell) {
		for _, cell := range row {
			cb(cell)
		}
	})
}

func (g Grid) Clone() *Grid {
	return NewGrid(g.dimensions.rows, g.dimensions.cols)
}
