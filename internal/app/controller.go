package app

import (
	"fmt"
)

var model *maze

var directions = map[string]string{"l": "left", "r": "right", "u": "up", "d": "down"}

// Start begins a new maze puzzle session.
func Start() {

	model = newMaze(4, 4)

	fmt.Println("Help me through the maze!")
	fmt.Println("Tell me which way to go - Left (l), Right (r), Up (u), or Down (d).")
	fmt.Printf("Hint: the maze is %v spaces wide by %v spaces high!\n", model.width, model.height)

	for !model.atEnd() {
		move()
	}
	fmt.Println("You made it out!")
	return
}

func move() {
	fmt.Printf("I am at space %v, %v.\n", model.location.Col+1, model.location.Row+1)
	dir := prompt()

	fmt.Printf("Moving %v\n", dir)
}

func prompt() string {
	var input string
	fmt.Println("which way?")
	fmt.Scanln(&input)
	if dir, ok := directions[input]; ok {
		return dir
	}
	fmt.Print("sorry, ")
	return prompt()
}
