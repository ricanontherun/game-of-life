package main

import (
	"fmt"
	"github.com/ricanontherun/game-of-life/src/grid"
	"time"
)

func printGrid(g *grid.Grid) {
	g.IterateRows(func(index int, row []grid.GridCell) {
		for colI := range row {
			fmt.Print(fmt.Sprintf("%d", g.GetCell(grid.GridCellPosition{
				Row: index,
				Col: colI,
			})))

			if colI != len(row) {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	})
}

func evolve(g grid.Grid) *grid.Grid {
	newGrid := g.Clone()

	// Generate new population.
	g.IterateCells(func(cell grid.GridCell) {
		evolvedCell := cell

		newGrid.SetCell(evolvedCell)
	})

	return newGrid
}

func main() {
	g := grid.NewGrid(5, 10)

	generation := 0
	for _ = range time.Tick(time.Second) {
		generation += 1

		fmt.Printf("Generation: %d\n", generation)

		printGrid(evolve(*g))

		break
	}
}
