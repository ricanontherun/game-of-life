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
func (grid Grid) Initialize(table [][]cellType) {
	assert(len(table) == grid.dimensions.rows, "Invalid Initialize size (rows)")
	assert(len(table[0]) == grid.dimensions.cols, "Invalid Initialize size (cols)")

	for rowI, row := range table {
		for colI, _ := range row {
			grid.SetCell(GridCell{
				Value: table[rowI][colI],
				Pos:   GridCellPosition{rowI, colI},
			})
		}
	}
}

func (grid Grid) GetCell(pos GridCellPosition) GridCell {
	assertRowBoundary(grid, pos.Row)
	assertColBoundary(grid, pos.Col)

	return grid.grid[pos.Row][pos.Col]
}

func (grid Grid) SetCell(cell GridCell) {
	assertRowBoundary(grid, cell.Pos.Row)
	assertColBoundary(grid, cell.Pos.Col)

	grid.grid[cell.Pos.Row][cell.Pos.Col] = cell
}

func (grid Grid) GetNeighbors(cell GridCell) map[string]GridCell {
	neighbors := make(map[string]GridCell)

	pos := cell.Pos

	neighbors["tl"] = grid.getRelativeNeighbor(pos, "tl")
	neighbors["tc"] = grid.getRelativeNeighbor(pos, "tc")
	neighbors["tr"] = grid.getRelativeNeighbor(pos, "tr")
	neighbors["l"] = grid.getRelativeNeighbor(pos, "l")
	neighbors["r"] = grid.getRelativeNeighbor(pos, "r")
	neighbors["bl"] = grid.getRelativeNeighbor(pos, "bl")
	neighbors["bc"] = grid.getRelativeNeighbor(pos, "bc")
	neighbors["br"] = grid.getRelativeNeighbor(pos, "br")

	return neighbors
}

func (grid Grid) GetLiveNeighborCount(cell GridCell) int {
	liveNeighbors := 0

	for _, neighbor := range grid.GetNeighbors(cell) {
		if neighbor.IsAlive() {
			liveNeighbors++
		}
	}

	return liveNeighbors
}

func (grid Grid) getRelativeNeighbor(pos GridCellPosition, deltaLabel string) GridCell {
	delta, exists := deltaMap[deltaLabel]
	if !exists {
		panic(errors.New("bad delta"))
	}

	neighborPosition := pos

	// Adjust the neighbors position on the grid.
	neighborPosition.Row += delta.deltaRow
	neighborPosition.Col += delta.deltaCol

	// Wrap any values which have left the board.
	if neighborPosition.Row < grid.boundaries.top {
		neighborPosition.Row = grid.boundaries.bottom
	} else if neighborPosition.Row > grid.boundaries.bottom {
		neighborPosition.Row = grid.boundaries.top
	}

	if neighborPosition.Col < grid.boundaries.left {
		neighborPosition.Col = grid.boundaries.right
	} else if neighborPosition.Col > grid.boundaries.right {
		neighborPosition.Col = grid.boundaries.left
	}

	return grid.GetCell(neighborPosition)
}

func (grid Grid) IterateRows(cb func(index int, row []GridCell)) {
	for index, row := range grid.grid {
		cb(index, row)
	}
}

func (grid Grid) IterateCells(cb func(cell GridCell)) {
	grid.IterateRows(func(rowIndex int, row []GridCell) {
		for _, cell := range row {
			cb(cell)
		}
	})
}

func (grid Grid) Clone() *Grid {
	return NewGrid(grid.dimensions.rows, grid.dimensions.cols)
}
