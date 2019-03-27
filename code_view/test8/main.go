package main

import "fmt"

type Cat struct {
	Name  string
	Color string
}

func (c *Cat) Meow() {
	fmt.Println("Name:", c.Name, "Color:", c.Color)
}

func main() {
	a := Cat{Name: "cat", Color: "black"}
	fmt.Println(a)

	b := &Cat{Name: "cat", Color: "black"}
	fmt.Println(b)
}
