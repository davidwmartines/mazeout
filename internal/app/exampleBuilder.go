package app

import g "github.com/davidwmartines/mazeout/internal/grid"

func exampleBuilder(maze *maze) {

	grid := maze.grid

	if len(grid) != 4 || len(grid[0]) != 4 {
		panic("example builder only works with 4x4 grid")
	}

	maze.destination = grid[2][3]
	maze.start = grid[1][0]

	walls := []*g.Point{}

	//top
	walls = append(walls, grid[0][0])
	walls = append(walls, grid[0][1])
	walls = append(walls, grid[0][2])
	walls = append(walls, grid[0][3])

	//bottom
	walls = append(walls, grid[3][0])
	walls = append(walls, grid[3][1])
	walls = append(walls, grid[3][2])
	walls = append(walls, grid[3][3])

	//left
	walls = append(walls, grid[2][0])

	//right
	walls = append(walls, grid[1][3])

	//obstacles
	walls = append(walls, grid[1][2])

	maze.walls = make(map[*g.Point]*g.Point)
	for _, wall := range walls {
		maze.walls[wall] = wall
	}
}
