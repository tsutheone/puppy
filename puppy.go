package puppy

import (
	"fmt"

	"github.com/tsutheone/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.DoggoGrowsUp(Bark())
}

func BigBarks() string {
	return dog.DoggoGrowsUp(Barks())
}

func From12() {
	fmt.Println("I'm from version v0.2.0")
}
