package main

import (
	"fmt"
)

/**
函数，简单来讲就是一段将输入数据转换为输出数据的公用代码块。当然有的时候函数的返回值为空，那么就是说输出数据为空。而真正的处理过程在函数内部已经完成了。

想一想我们为什么需要函数，最直接的需求就是代码中有太多的重复代码了，为了代码的可读性和可维护性，将这些重复代码重构为函数也是必要的。


*/
func main() {
	//函数定义
	function_define()
	//多返回值
	mul_return()
	//变长参数
	can_change_param_len()
	//闭包函数
	close_package_function()
	//递归函数
	recursion_function()
	//异常处理
	exception_handle()
}

/**
异常处理

当你读取文件失败而退出的时候是否担心文件句柄是否已经关闭？
抑或是你对于try…catch…finally的结构中finally里面的代码和try里面的return代码那个先执行这样的问题痛苦不已？

一切都结束了。一门完美的语言必须有一个清晰的无歧义的执行逻辑。
*/
func exception_handle() {
	exception_handle_test1()
	exception_handle_test3()
	exception_handle_test2()
}

/**
Go语言提供了关键字defer来在函数运行结束的时候运行一段代码或调用一个清理函数。
下面的例子中，虽然second()函数写在first()函数前面，但是由于使用了defer标注，所以它是在main函数执行结束的时候才调用的。

例如：
使用os包中的Open方法打开文件后，立马跟着一个defer语句用来关闭文件句柄。
这样就保证了该文件句柄在main函数运行结束的时候或者异常终止的时候一定能够被释放。
而且由于紧跟着Open语句，一旦养成了习惯，就不会忘记去关闭文件句柄了。
*/
func exception_handle_test1() {
	fmt.Println("======defer 功能简单介绍")
	defer second()
	first()
}

/**
panic & recover

当你周末走在林荫道上，听着小歌，哼着小曲，很是惬意。
突然之间，从天而降瓢泼大雨，你顿时慌张（panic）起来，没有带伞啊，淋着雨感冒就不好了。
于是你四下张望，忽然发现自己离地铁站很近，那里有很多卖伞的，心中顿时又安定了下来（recover），于是你飞奔过去买了一把伞（defer）。

好了，panic和recover是Go语言提供的用以处理异常的关键字。
panic用来触发异常，而recover用来终止异常并且返回传递给panic的值。
（注意recover并不能处理异常，而且recover只能在defer里面使用，否则无效。）
*/
func exception_handle_test2() {
	fmt.Println("=====panic & recover的特性")

	fmt.Println("I am walking and singing...")
	//当panic时程序就已经终于了
	panic("It starts to rain cats and dogs")
	//所以下面的语句不会被执行
	msg := recover()
	fmt.Println(msg)
}

func exception_handle_test3() {
	fmt.Println("======解决panic后不执行的问题")
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	fmt.Println("I am walking and singing...")
	//当panic时程序就已经终于了
	panic("It starts to rain cats and dogs")
}

/**

 */
func first() {
	fmt.Println("the first func")
}

func second() {
	fmt.Println("the second func")
}

/**
每次谈到递归函数，必然绕不开阶乘和斐波拉切数列。
*/
func recursion_function() {
	fmt.Println("-----递归-------------------------")
	factorial_test()
	fibonacci_test()
}

/**
阶乘
*/
func factorial_test() {
	fmt.Println("=====阶乘")
	fmt.Println(factorial(5))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

/**
斐波拉切数列
*/
func fibonacci_test() {
	fmt.Println("=====斐波拉切数列")
	fmt.Println(fibonacci(8))
}

func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return fibonacci(n-2) + fibonacci(n-1)
	}
}

/**
闭包函数
*/
func close_package_function() {
	fmt.Println("====闭包函数 闭包函数对它外层的函数中的变量具有访问和修改的权限")
	base := 20
	arr := []int{1, 2, 3, 4, 5}
	sum := func(arr ...int) int {
		sum_ := 0
		for i := 0; i < len(arr); i++ {
			sum_ += arr[i]
		}
		return sum_ + base
	}
	fmt.Println(sum(arr...))

	//闭包的示例
	fmt.Println("=====闭包示例 生成偶数序列")

	/**
	  重点在于，返回的值是函数
	  return func() (retVal uint) {
	          retVal = i
	          i += 2
	          return
	      }
	  所以每次调用nextEven时，只执行返回的函数， createEvenGenerator()函数中的i := uint(0)只会被执行一次。
	  由于闭包函数可以使用外层函数的变量，所以每调用一次nextEven函数时，外层函数中的i都会被+2
	*/
	nextEven := createEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

}

func createEvenGenerator() func() uint {
	i := uint(0)
	fmt.Println("i=", i)
	return func() (retVal uint) {
		retVal = i
		i += 2
		return
	}
}

/**
可变长度参数
*/
func can_change_param_len() {
	fmt.Println("====变长参数 可变长参数定义只能是函数的最后一个参数")
	fmt.Println(sum_handle(1, 2, 3, 4))
}

func sum_handle(arr ...int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}

	return sum
}

/**
多返回值

这是GO灵活的地方，JAVA中还需要定义实体对象
*/
func mul_return() {

	fmt.Println("=====多返回值")
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr_sum_avg(arr))
}

func arr_sum_avg(arr []int) (int, float64) {
	sum := 0
	avg := 0.0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	avg = float64(sum) / float64(len(arr))

	return sum, avg
}

/**
函数定义

预定义返回值

实参数与虚参数
实参数：就是函数调用的时候传入的参数
虚参数：就是函数定义的时候表示函数需要传入哪些参数的占位参数
*/
func function_define() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{10, 11, 12, 13, 14, 15}

	fmt.Println(slice_sum(arr1))
	fmt.Println(slice_sum_predefine(arr2))
}

/**
标准函数
*/
func slice_sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

/**
预先定义返回值，函数结束后直接return就可以了
(sum int) 就是预定义返回值，并且函数体中无需再定义sum
*/
func slice_sum_predefine(arr []int) (sum int) {
	sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return
}
