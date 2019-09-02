package app

import (
	g "github.com/davidwmartines/mazeout/internal/grid"
)

type maze struct {
	width       int
	height      int
	grid        g.Grid
	walls       map[*g.Point]*g.Point
	location    *g.Point
	destination *g.Point
	start       *g.Point
}

type builder func(maze *maze)

type direction string

const left direction = "left"
const right direction = "right"
const up direction = "up"
const down direction = "down"

func newMaze(height, width int, builderName string) *maze {

	maze := maze{}
	maze.width = width
	maze.height = height
	maze.grid = g.Create(height, width)

	build := makeBuilder(builderName)
	build(&maze)

	maze.location = maze.start

	return &maze
}

func (maze *maze) isValid(dir direction) bool {
	next := getNext(maze, dir)
	return next != nil
}

func (maze *maze) isWall(dir direction) bool {
	next := getNext(maze, dir)
	if next == nil {
		panic("invalid direction")
	}
	_, contains := maze.walls[next]
	return contains
}

func getNext(maze *maze, dir direction) *g.Point {
	loc := maze.location
	neighbors := maze.grid.Neighbors(loc)

	var next *g.Point
	for _, n := range neighbors {
		if dir == left {
			if n.Row == loc.Row && n.Col == loc.Col-1 {
				next = n
				break
			}
		}
		if dir == right {
			if n.Row == loc.Row && n.Col == loc.Col+1 {
				next = n
				break
			}
		}
		if dir == up {
			if n.Col == loc.Col && n.Row == loc.Row-1 {
				next = n
				break
			}
		}
		if dir == down {
			if n.Col == loc.Col && n.Row == loc.Row+1 {
				next = n
				break
			}
		}
	}
	return next
}

func (maze *maze) move(dir direction) {
	maze.location = getNext(maze, dir)
}

func (maze *maze) atEnd() bool {
	return maze.location == maze.destination
}
