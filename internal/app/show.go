package app

import "fmt"

// "cheat" function to show the maze on screen
func show(maze *maze) {

	hLine := " --"
	for range maze.grid[0] {
		hLine += "-"
	}
	fmt.Println(hLine)

	for _, row := range maze.grid {
		rowOut := " |"
		for _, point := range row {
			if _, iswall := maze.walls[point]; iswall {
				rowOut += "█"
			} else if point == maze.start {
				rowOut += "s"
			} else if point == maze.destination {
				rowOut += "d"
			} else {
				rowOut += " "
			}
		}
		rowOut += "|"
		fmt.Println(rowOut)
	}
	fmt.Println(hLine)
}
