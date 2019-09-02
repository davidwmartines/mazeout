package app

import "fmt"

func makeBuilder(builderName string) builder {

	switch builderName {

	case "example":
		return exampleBuilder
	case "life":
		return lifeBuilder
	default:
		panic(fmt.Sprintf("unknown builder name: '%v'", builderName))
	}
}
