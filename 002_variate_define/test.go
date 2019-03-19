package main

import "fmt"

func main() {
	fmt.Println("------变量------------------")
	variate()
	fmt.Println("------常量------------------")
	constant()
}

/**
常量
Go语言也支持常量定义。所谓常量就是在程序运行过程中保持值不变的变量定义。
常量的定义和变量类似，只是用const关键字替换了var关键字，另外常量在定义的时候必须有初始值。
*/
func constant() {
	fmt.Println("第一式：标准式")
	const x string = "hello world"
	fmt.Println(x)
	fmt.Println("第二式：不指定类型，由系统自动推断")
	const y = "hello world"
	fmt.Println(y)
}

/**
变量之所以称为变量，就是因为它们的值在程序运行过程中可以发生变化，但是它们的变量类型是无法改变的。因为Go语言是静态语言，
并不支持程序运行过程中变量类型发生变化。比如如果你强行将一个字符串值赋值给定义为int的变量，那么会发生编译错误。
即使是强制类型转换也是不可以的。强制类型转换只支持同类的变量类型。比如数值类型之间强制转换。
*/
func variate() {
	test1()
	test2()
	test3()
}

func test1() {
	fmt.Println("第一式： 标准式")
	var a string = "helloworld"
	fmt.Println(a)
}

func test2() {
	fmt.Println("第二式：预定义式")
	var a string
	a = "helloworld"
	fmt.Println(a)
}

func test3() {
	fmt.Println("第三式：不指定类型，由系统推断")
	var a = "helloworld"
	fmt.Println(a)
}

func test4() {
	fmt.Println("第四式：简洁式，连var都可以省略，但只能在函数内部使用")
	a := "helloworld"
	fmt.Println(a)
}
