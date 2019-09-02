package app

import "fmt"

// "cheat" function to show the maze on screen
func show(maze *maze) {
	for _, row := range maze.grid {
		rowOut := " "
		for _, point := range row {
			if _, iswall := maze.walls[point]; iswall {
				rowOut += "█"
			} else if point == maze.start {
				rowOut += " "
			} else if point == maze.destination {
				rowOut += "*"
			} else {
				rowOut += " "
			}
		}
		fmt.Println(rowOut)
	}
}
