package puppy

import (
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
