package main

import "fmt"

func main() {
	do1()
	do2()
}

func do1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println("do1() ", s)
}

func do2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println("do2() ", s) //[1 2 3]
}
