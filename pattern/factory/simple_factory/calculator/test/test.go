package main

import (
	"fmt"
	"go_study/pattern/factory/simple_factory/calculator"
)

func main() {
	a := 9
	b := 3
	result, err := calculator.ToCalulator(a, b, "&")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
