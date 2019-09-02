package app

import (
	"fmt"
)

var model *maze

const maxHelps = 10

var helpsUsed int

const help direction = "help"

var directions = map[string]direction{"l": left, "r": right, "u": up, "d": down, "h": help}

// Start begins a new maze puzzle session.
func Start() {

	//model = newMaze(4, 4, "example")
	model = newMaze(10, 30, "life")

	fmt.Println("Help me through the maze!")
	fmt.Printf("The maze is a 0-indexed grid of %v rows by %v columns, with %v wall points.\n", model.height, model.width, len(model.walls))
	fmt.Printf("I am starting at row %v, column %v, and I need to get to row %v, column %v.\n", model.start.Row, model.start.Col, model.destination.Row, model.destination.Col)
	fmt.Println("Tell me which way to go - Left (l), Right (r), Up (u), or Down (d).")
	fmt.Printf("Hint: You can also ask for help (h) to view the maze, up to %v times.\n", maxHelps)
	fmt.Println("")

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
	if dir == help {
		show(model)
		helpsUsed++
		fmt.Printf("You have used %v of %v available helps.\n", helpsUsed, maxHelps)
		move()
		return
	}
	fmt.Printf("Moving %v\n", dir)
	model.move(dir)
}

func isValid(dir direction) bool {
	if dir == help {
		if helpsUsed < maxHelps {
			return true
		}
		fmt.Println("Sorry, you have used all your help!")
		return false
	}
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
