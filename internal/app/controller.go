package app

import (
	"fmt"
)

var model *maze

var directions = map[string]direction{"l": left, "r": right, "u": up, "d": down}

// Start begins a new maze puzzle session.
func Start() {

	model = newMaze(4, 4, "example")

	fmt.Println("Help me through the maze!")
	fmt.Println("Tell me which way to go - Left (l), Right (r), Up (u), or Down (d).")
	fmt.Printf("Hint: the maze is a 0-indexed grid of %v rows by %v columns!\n", model.height, model.width)

	for !model.atEnd() {
		move()
	}
	fmt.Println("Yay! I made it out. Thank you!")
	return
}

func move() {
	fmt.Printf("I am at row %v, column %v.\n", model.location.Row, model.location.Col)
	var dir direction
	var valid = false
	for !valid {
		dir = prompt()
		valid = isValid(dir)
	}
	fmt.Printf("Moving %v\n", dir)
	model.move(dir)
}

func isValid(dir direction) bool {
	if !model.isValid(dir) {
		fmt.Println("Can't go that way.")
		return false
	}
	if model.isWall(dir) {
		fmt.Printf("Can't go %v into wall!\n", dir)
		return false
	}
	return true
}

func prompt() direction {
	var input string
	fmt.Print("which way? ")
	fmt.Scanln(&input)
	if dir, ok := directions[input]; ok {
		return dir
	}
	fmt.Print("sorry, ")
	return prompt()
}
