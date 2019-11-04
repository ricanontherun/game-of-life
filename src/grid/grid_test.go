package grid

import (
	"math/rand"
	"testing"
)

func getRandomTable(rows int, cols int) [][]cellType {
	table := make([][]cellType, rows)

	for i := range table {
		row := make([]cellType, cols)

		for i := 0; i < cols; i++ {
			row[i] = cellType(rand.Int())
		}

		table[i] = row
	}

	return table
}

// Test object initializations
func TestGridInitialization(t *testing.T) {
	rows := 8
	cols := 10

	g := NewGrid(rows, cols)

	if g.dimensions.cols != cols {
		t.Error("Dimensions.cols initialized incorrectly")
	}

	if g.dimensions.rows != rows {
		t.Error("Dimensions.rows initialized incorrectly")
	}

	if g.boundaries.bottom != (rows - 1) {
		t.Error("Dimensions.boundaries initialized incorrectly")
	}

	if g.boundaries.top != 0 {
		t.Error("Dimensions.boundaries initialized incorrectly")
	}

	if g.boundaries.left != 0 {
		t.Error("Dimensions.boundaries initialized incorrectly")
	}

	if g.boundaries.right != (cols - 1) {
		t.Error("Dimensions.boundaries initialized incorrectly")
	}
}

func TestSetCell(t *testing.T) {
	g := NewGrid(10, 10)

	g.SetCell(GridCell{
		13,
		GridCellPosition{
			2,
			1,
		},
	})

	if g.grid[2][1].Value != 13 {
		t.Error("Failed to set cell value")
	}
}

func TestGetCell(t *testing.T) {
	g := NewGrid(10, 10)

	// Prop up some data
	g.grid[9][2].Value = 100

	if g.GetCell(GridCellPosition{
		Row: 9,
		Col: 2,
	}).Value != 100 {
		t.Error("Cell value should have been 100")
	}

	if g.GetCell(GridCellPosition{
		Row: 0,
		Col: 0,
	}).Value != 0 {
		t.Error("Cell value should have been 0")
	}
}

func TestGetRelativeNeighbor(t *testing.T) {
	g := NewGrid(5, 5)
	randomTable := getRandomTable(5, 5)
	g.initialize(randomTable)

	type testRow struct {
		centralRow  int
		centralCol  int
		deltaLabel  string
		expectedRow int
		expectedCol int
	}

	tests := []testRow{
		{1, 1, "tl", 0, 0},
		{0, 0, "tl", 4, 4},
		{1, 1, "tc", 0, 1},
		{0, 1, "tc", 4, 1},
		{2, 4, "tc", 1, 4},
		{3, 3, "tr", 2, 4},
		{4, 4, "l", 4, 3},
		{0, 3, "l", 0, 2}, // wrap
		{0, 1, "r", 0, 2},
		{0, 4, "r", 0, 0}, // wrap
		{4, 4, "bl", 0, 3},
		{4, 4, "bc", 0, 4},
		{4, 4, "br", 0, 0},
		{3, 2, "br", 4, 3},
	}

	for _, test := range tests {
		neighbor := g.getRelativeNeighbor(GridCellPosition{
			Row: test.centralRow,
			Col: test.centralCol,
		}, test.deltaLabel)

		if neighbor.Value != g.grid[test.expectedRow][test.expectedCol].Value {
			t.Errorf("BAN")
		}
	}
}
