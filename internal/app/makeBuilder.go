package app

import "fmt"

func makeBuilder(builderName string) builder {

	switch builderName {

	case "example":
		return exampleBuilder
	default:
		panic(fmt.Sprintf("unknown builder name: '%v'", builderName))
	}
}
