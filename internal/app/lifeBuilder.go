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
				if len(maze.grid.Neighbors(point)) > 3 {
					maze.start = point
					break
				}
			}
		}
		if maze.start != nil {
			break
		}
	}

	//set destination to first non-wall point && non-corner, working backwards
	for rowIndex := len(maze.grid) - 1; rowIndex >= 0; rowIndex-- {
		row := maze.grid[rowIndex]
		for pointIndex := len(row) - 1; pointIndex >= 0; pointIndex-- {
			point := row[pointIndex]
			if _, iswall := maze.walls[point]; !iswall {
				if len(maze.grid.Neighbors(point)) > 3 {
					maze.destination = point
					break
				}
			}
		}
		if maze.destination != nil {
			break
		}
	}

	//add boundary walls
	//top
	for _, point := range maze.grid[0] {
		if point != maze.start && point != maze.destination {
			maze.walls[point] = point
		}
	}
	//bottom
	for _, point := range maze.grid[len(maze.grid)-1] {
		if point != maze.start && point != maze.destination {
			maze.walls[point] = point
		}
	}
	//left/right
	for _, row := range maze.grid {
		leftPoint := row[0]
		if leftPoint != maze.start && leftPoint != maze.destination {
			maze.walls[leftPoint] = leftPoint
		}
		rightPoint := row[len(row)-1]
		if rightPoint != maze.start && rightPoint != maze.destination {
			maze.walls[rightPoint] = rightPoint
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
	max := 55
	return generator.Intn(max-min+1) + min

}
