package main

import "fmt"

func adder1() func(i int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func doAdder1() {
	//a := adder1()表达式应该是得到下面这个函数
	//return func(v int) int {
	//	sum += v
	//	return sum
	//}
	a := adder1()

	//每次往a函数中传入参数都会得到自加
	//sum := 0 -> sum是个函数闭包内的自由变量
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}

type iAdder func(i int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

//这是正统的函数式编程
func doAdder2() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0+1+...%d = %d\n", i, s)
	}
}

func main() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0+1+...%d = %d\n", i, s)
	}
}
