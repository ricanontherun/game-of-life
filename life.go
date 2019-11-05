package main

import (
	"fmt"
	"github.com/ricanontherun/game-of-life/src/grid"
	"time"
)

type GameBoard struct {
	g          *grid.Grid
	generation int
}

func printGrid(g *grid.Grid) {
	g.IterateRows(func(index int, row []grid.GridCell) {
		for colI, cell := range row {
			fmt.Print(fmt.Sprintf("%d", cell.Value))

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
	board := GameBoard{
		g:          grid.NewGrid(10, 10),
		generation: 0,
	}

	// Setup an initial configuration.
	board.g.SetCell(grid.GridCell{
		Value: 1,
		Pos: grid.GridCellPosition{
			Row: 3,
			Col: 5,
		},
	})

	board.g.SetCell(grid.GridCell{
		Value: 1,
		Pos: grid.GridCellPosition{
			Row: 4,
			Col: 5,
		},
	})

	board.g.SetCell(grid.GridCell{
		Value: 1,
		Pos: grid.GridCellPosition{
			Row: 5,
			Col: 5,
		},
	})

	board.generation = 1
	fmt.Printf("Generation: %d\n", board.generation)
	printGrid(board.g)

	for _ = range time.Tick(time.Second) {
		board.generation += 1
		fmt.Printf("Generation: %d\n", board.generation)
		board.g = evolve(*board.g)
		printGrid(board.g)

		break
	}
}
