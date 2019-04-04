package strategy

import (
	"fmt"
	"testing"
	"time"
)

func testStrategy(t *testing.T) {
	time.Sleep(time.Second)
	cashContext := NewCashContext("normal")
	result := cashContext.AcceptCash(98.9)

	fmt.Println(result)
}
