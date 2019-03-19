package main

import (
	"fmt"
	"go_study/007_struct_interface/animal/animal_type/dog"
	"go_study/007_struct_interface/animal/animal_type/fish"
	"go_study/007_struct_interface/animal/country/china"
	"go_study/007_struct_interface/animal/country/english"
)

type Animal interface {
	fish.Fish
	dog.Dog
}

func DoAllAnimal(animal Animal) {
	switch v := animal.(type) {
	case china.AnimalImpl:
		fmt.Println(v.ActionType2)
	case english.AnimalImpl:
		fmt.Println(v.ActionType1)
	}

	fmt.Println(animal.Barks())
	fmt.Println(animal.Swimming())

}

func main() {
	animal := china.AnimalImpl{"游啊游啊游。。。。", "汪汪汪。。。", "一群含蓄的动物们："}
	DoAllAnimal(animal)

	animal = china.AnimalImpl{"使劲的游啊游啊游。。。。", "使劲的汪汪汪。。。", "一群打了鸡血的动物们："}
	DoAllAnimal(animal)
}
