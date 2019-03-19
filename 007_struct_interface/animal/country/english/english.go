package english

import (
	"fmt"
)

type AnimalImpl struct {
	SwimingStr  string
	BarkStr     string
	ActionType1 string
}

func (animal AnimalImpl) Swimming() string {
	return animal.SwimingStr
}

func (animal AnimalImpl) Barks() string {
	return animal.BarkStr
}

func (animal AnimalImpl) ActionType() {
	fmt.Println(animal.ActionType1)
}
