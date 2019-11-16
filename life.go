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
		aliveCount := g.GetLiveNeighborCount(cell)

		if cell.IsAlive() { // Live actions
			if aliveCount < 2 {
				// Any live cell with fewer than two live neighbors dies
				evolvedCell.Kill()
			} else if aliveCount > 3 {
				// Any live cell with more than 3 neighbors dies
				evolvedCell.Kill()
			}
		} else {
			if aliveCount == 3 {
				evolvedCell.Resurrect()
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
			Row: 3,
			Col: 6,
		},
	})
	board.g.SetCell(grid.GridCell{
		Value: 1,
		Pos: grid.GridCellPosition{
			Row: 4,
			Col: 6,
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
	}
}
