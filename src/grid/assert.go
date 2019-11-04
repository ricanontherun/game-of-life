package grid

import (
	"errors"
)

func assert(cond bool, msg string) {
	if !cond {
		panic(errors.New(msg))
	}
}

func assertRowBoundary(grid Grid, row int) {
	assert(row >= grid.boundaries.top && row <= grid.boundaries.bottom, "Out of bounds (row)")
}

func assertColBoundary(grid Grid, col int) {
	assert(col >= grid.boundaries.left && col <= grid.boundaries.right, "Out of bounds (col)")
}
