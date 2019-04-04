package main

import (
	"fmt"
	"go_study/pattern/strategy"
)

func main() {
	cashContext := strategy.NewCashContext("normal")
	result := cashContext.AcceptCash(905.2)
	fmt.Println(result)

	cashContext = strategy.NewCashContext("300减100")
	result = cashContext.AcceptCash(905.2)
	fmt.Println(result)

	cashContext = strategy.NewCashContext("打八折")
	result = cashContext.AcceptCash(905.2)
	fmt.Println(result)
}
