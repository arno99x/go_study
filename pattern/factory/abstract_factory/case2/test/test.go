package main

import "go_study/pattern/factory/abstract_factory/case2"

func main() {
	factory.FastFoodFactoryA{}.CreateFood().Eat()
	factory.FastFoodFactoryA{}.CreateDrink().Drink()

	factory.FastFoodFactoryB{}.CreateFood().Eat()
	factory.FastFoodFactoryB{}.CreateDrink().Drink()

}
