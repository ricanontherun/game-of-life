package game

import "github.com/ricanontherun/conway/src/grid"

type Board struct {
	grid.Grid
}

func (b Board) evolve() Board {
	// Apply game of life rules.
}