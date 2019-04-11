package main

import (
	"fmt"
	"go_study/base/007_struct_interfacerface/duck/green"
	"go_study/base/007_struct_interfacerface/duck/yellow"
)

type Duck interface {
	GetInfo() string
}

func getDuck(duck Duck) string {
	return duck.GetInfo()
}

func main() {
	var duck Duck
	duck = &green.Duck{"green", 5.3, "value2.1"}
	fmt.Printf("type -> %T , v->%v\n", duck, duck)
	inpect(duck)

	duck = yellow.Duck{"yellow", 3.2, "value2.2"}
	fmt.Printf("type -> %T , v->%v\n", duck, duck)
	inpect(duck)

}

func inpect(duck Duck) {
	switch v := duck.(type) {
	case yellow.Duck:
		fmt.Println(v.YellowValue)
	case *green.Duck:
		fmt.Println(v.GreenValue)

	}
}
