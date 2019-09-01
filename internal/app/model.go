package app

import (
	g "github.com/davidwmartines/mazeout/internal/grid"
)

type maze struct {
	width       int
	height      int
	grid        g.Grid
	walls       []*g.Point
	location    *g.Point
	destination *g.Point
}

func (maze *maze) atEnd() bool {
	return maze.location == maze.destination
}

func newMaze(height, width int) *maze {

	maze := maze{}
	maze.width = width
	maze.height = height
	maze.grid = g.Create(height, width)

	maze.walls = buildWalls(maze.grid)
	maze.destination = findExit(&maze)
	maze.location = findStart(&maze)
	return &maze
}

func buildWalls(grid g.Grid) (walls []*g.Point) {

	walls = append(walls, grid[0][1])
	walls = append(walls, grid[0][2])
	walls = append(walls, grid[2][1])
	walls = append(walls, grid[2][2])
	return
}

func findExit(maze *maze) *g.Point {
	lastRow := maze.grid[len(maze.grid)-1]
	return lastRow[len(lastRow)-1]
}

func findStart(maze *maze) *g.Point {
	return maze.grid[0][0]
}
