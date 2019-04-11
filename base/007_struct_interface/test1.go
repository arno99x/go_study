package main

import "fmt"

func main() {
	struct_test()
	interface_test()
}

/**
结构体定义

上面我们说过Go的指针和C的不同，结构体也是一样的。Go是一门删繁就简的语言，一切令人困惑的特性都必须去掉。

简单来讲，Go提供的结构体就是把使用各种数据类型定义的不同变量组合起来的高级数据类型。
*/
func struct_test() {
	fmt.Println("------结构体-------------------------")
	struct_test1()
	struct_test2()
	struct_test3()
	struct_test4()
	struct_test5()
	fmt.Println("------接口---------------------------")
	interface_test()
}

/**
定义矩形结构体

或者这样写

type Rect2 struct {
    width,length float64
}
*/

type Rect struct {
	width  float64
	length float64
}

/**
结构体的创建
*/
func struct_test1() {
	fmt.Println("======结构体的创建")
	var rect Rect
	rect.width = 10
	rect.length = 15

	//或者这样也行
	rect1 := Rect{width: 11, length: 21}

	//如果你知道结构体的顺序也可以这样赋值
	rect2 := Rect{50, 100}

	fmt.Println(rect.width * rect.length)
	fmt.Println(rect1.width * rect1.length)
	fmt.Println("rect2.width = ", rect2.width, "   rect2.length = ", rect2.length)

}

/**
GO中的参数传递方式是值传递，对于结构体也同要是值传递
*/
func struct_test2() {
	fmt.Println("=====GO中的参数传递方式是值传递，对于结构体也同要是值传递")
	rect := Rect{5, 10}

	fmt.Println(double_area(rect))

	fmt.Println(rect.length, rect.width)
}

func double_area(rect Rect) float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

/**
结构体组合函数

咦！。。。 area()并没有定义在Rect struct结构体中呀，为什么会变成Rect的函数呢？

关键字func表示这是一个函数
第二个参数是结构体类型和实例变量
第三个是函数名称
第四个是函数返回值。

关键是第二个参数
这里我们可以看出area()函数和普通函数定义的区别就在于area()函数多了一个结构体类型限定。这样一来Go就知道了这是一个为结构体定义的方法。
*/
func struct_test3() {
	rect := Rect{20, 40}
	fmt.Println(rect.area())
}

func (rect Rect) area() float64 {
	return rect.width * rect.length
}

/**
结构体内嵌类型

我们可以在一个结构体内部定义另外一个结构体类型的成员。例如iPhone也是Phone，我们看下例子：
*/
func struct_test4() {
	fmt.Println("======结构体内嵌类型")
	var p IPhone
	p.phone.color = "red"
	p.phone.price = 6000
	p.modle = "iphone"

	//这种给人的感觉是 ， phone属于iphone ，但这里想表达的是iphone是phone的一种
	fmt.Println("I have a iPhone:")
	fmt.Println("Price:", p.phone.price)
	fmt.Println("Color:", p.phone.color)
	fmt.Println("Model:", p.modle)

	//所以我们可以这样做，换一种方式定义结构体
	var p1 IPhone1
	p1.color = "blu"
	p1.price = 7000
	p1.modle = "iphone x"
	fmt.Println("I have other iPhone:")
	fmt.Println("Price:", p1.price)
	fmt.Println("Color:", p1.color)
	fmt.Println("Model:", p1.modle)

}

type Phone1 struct {
	price float64
	color string
}

type IPhone1 struct {
	Phone
	modle string
}

type Phone struct {
	price float64
	color string
}

type IPhone struct {
	phone Phone
	modle string
}

/**
如果结构体A内嵌了结构体B，那么A也可以调用B的组合函数
*/
func struct_test5() {
	var p1 IPhone1
	p1.ringing()
}

func (phone1 Phone) ringing() {
	fmt.Println("Phone is ringing...")
}

func interface_test() {
	var phones = [5]Phone2{
		NokiaPhone2{price: 350},
		IPhone2{price: 5000},
		IPhone2{price: 3400},
		NokiaPhone2{price: 450},
		IPhone2{price: 5000},
	}
	var totalSales = 0
	for _, phone := range phones {
		totalSales += phone.sales()
	}
	fmt.Println(totalSales)
}

type Phone2 interface {
	call()
	sales() int
}
type NokiaPhone2 struct {
	price int
}

func (nokiaPhone NokiaPhone2) call() {
	fmt.Println("I am Nokia, I can call you!")
}
func (nokiaPhone NokiaPhone2) sales() int {
	return nokiaPhone.price
}

type IPhone2 struct {
	price int
}

func (iPhone IPhone2) call() {
	fmt.Println("I am iPhone, I can call you!")
}
func (iPhone IPhone2) sales() int {
	return iPhone.price
}
