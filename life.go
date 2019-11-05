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

		aliveCount := 0

		for _, neighbor := range g.GetNeighbors(cell) {
			if neighbor.Value == 1 {
				aliveCount++
			}

			if aliveCount == 4 { // 4 is the highest number of meaningful live neighbors
				break
			}
		}

		if cell.Value == 1 { // Live actions
			if aliveCount < 2 {
				evolvedCell.Value = 0
			} else if aliveCount > 3 {
				evolvedCell.Value = 0
			}
		} else {
			if aliveCount == 3 {
				evolvedCell.Value = 1
			}
		}

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
