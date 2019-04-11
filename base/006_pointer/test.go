package main

import "fmt"

/**
指针的一大用途就是可以将变量的指针作为实参传递给函数，从而在函数内部能够直接修改实参所指向的变量值。
*/
func main() {
	var x int
	var x_ptr *int
	x = 10
	x_ptr = &x

	/**
	    & 取一个变量的地址
	    * 取一个指针变量所指向的地址的值
	**/
	//输出x的值
	fmt.Println(x)
	//输出x的地址
	fmt.Println(x_ptr)
	fmt.Println("..", &x_ptr)
	fmt.Println("...", *&x_ptr)
	//输出x地址对应的x的值
	fmt.Println(*x_ptr)
}
