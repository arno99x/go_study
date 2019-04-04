package main

import "fmt"

type Component struct {
	Component   *Component
	DoOperation func()
}

func main() {
	component1 := Component{}
	component1.DoOperation = func() {
		fmt.Println("component1 do operation ... ")
	}

	component2 := Component{}
	component2.Component = &component1
	component2.DoOperation = func() {
		component2.Component.DoOperation()
		fmt.Println("component2 do operation ...")
	}

	component3 := Component{}
	component3.Component = &component2
	component3.DoOperation = func() {
		component3.Component.DoOperation()
		fmt.Println("component3 do operation ...")
	}

	component3.DoOperation()
}
