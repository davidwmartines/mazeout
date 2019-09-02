package app

import (
	"math/rand"
	"time"

	life "github.com/davidwmartines/life/pkg/model"
	"github.com/davidwmartines/life/pkg/seeds"
	gr "github.com/davidwmartines/mazeout/internal/grid"
)

// builds the maze using cellular animation (conway's game of life)
func lifeBuilder(maze *maze) {

	maze.walls = make(map[*gr.Point]*gr.Point)

	// run a life model
	lifeModel := life.Create(maze.height, maze.width)
	seed := pickSeed(maze)
	generations := pickGenerations(maze, seed)
	lifeModel.ApplySeed(seed)
	for i := 1; i <= generations; i++ {
		lifeModel.Tick()
	}

	// use alive cells as the walls
	for _, cell := range lifeModel.Cells {
		if cell.Alive {
			mazePoint := maze.grid[cell.Point.Row][cell.Point.Col]
			maze.walls[mazePoint] = mazePoint
		}
	}

	//set start to first non-wall point
	for _, row := range maze.grid {
		for _, point := range row {
			if _, iswall := maze.walls[point]; !iswall {
				maze.start = point
				break
			}
		}
		if maze.start != nil {
			break
		}
	}

	//set destination to first non-wall point working backwards
	for rowIndex := len(maze.grid) - 1; rowIndex >= 0; rowIndex-- {
		row := maze.grid[rowIndex]
		for pointIndex := len(row) - 1; pointIndex >= 0; pointIndex-- {
			point := row[pointIndex]
			if _, iswall := maze.walls[point]; !iswall {
				maze.destination = point
				break
			}
		}
		if maze.destination != nil {
			break
		}
	}
}

func pickSeed(maze *maze) life.Seed {
	//TODO: based on maze dimensions, use a good seed pattern.
	return seeds.Acorn
}

func pickGenerations(maze *maze, seed life.Seed) int {
	//TODO: based on maze dimensions and seed, calculate a good number of generations
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	min := 25
	max := 35
	return generator.Intn(max-min+1) + min

}
