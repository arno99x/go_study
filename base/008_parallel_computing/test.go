package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
并行计算是GO与生具来的特点
*/
func main() {
	test1()
	test2()
	test3()

	channel_test()
}

/**
Channel提供了协程之间的通信方式以及运行同步机制。

假设训练定点投篮和三分投篮，教练在计数。
*/
func channel_test() {
	fmt.Println("======Channel提供了协程之间的通信方式以及运行同步机制。")
	var c chan string
	c = make(chan string)
	go fixed_shooting(c)
	go count(c)
	var input string
	fmt.Scanln(&input)
}

func fixed_shooting(msg_chan chan string) {
	for {
		msg_chan <- "fixed shooting"
		fmt.Println("continue fixed shooting...")
	}
}
func count(msg_chan chan string) {
	for {
		msg := <-msg_chan
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

/**
做个题目

如果我们列出10以下所有能够被3或者5整除的自然数，那么我们得到的是3，5，6和9。这四个数的和是23。
那么请计算1000以下（不包括1000）的所有能够被3或者5整除的自然数的和。

这个题目的一个思路就是：

(1) 先计算1000以下所有能够被3整除的整数的和A，
(2) 然后计算1000以下所有能够被5整除的整数和B，
(3) 然后再计算1000以下所有能够被3和5整除的整数和C，
(4) 使用A+B-C就得到了最后的结果。

按照上面的方法，传统的方法当然就是一步一步计算，然后再到第(4)步汇总了。

但是一旦有了Go，我们就可以让前面三个步骤并行计算，然后再在第(4)步汇总。

并行计算涉及到一个新的数据类型chan和一个新的关键字go。
*/
func test1() {
	fmt.Println("实例1：范围内能被3或5整除的数之和")
	LIMIT := 10000
	resultChan := make(chan int, 3)
	t_start := time.Now()
	go get_sum_of_divisible(LIMIT, 3, resultChan)
	go get_sum_of_divisible(LIMIT, 5, resultChan)
	//这里其实那个是被3整除，哪个是被5整除看具体调度方法，不过由于是求和，所以没关系
	sum3, sum5 := <-resultChan, <-resultChan
	//单独算被15整除的
	go get_sum_of_divisible(LIMIT, 15, resultChan)
	sum15 := <-resultChan
	sum := sum3 + sum5 - sum15
	t_end := time.Now()
	fmt.Println(sum)
	fmt.Println(t_end.Sub(t_start))
}

func get_sum_of_divisible(num int, divider int, resultChan chan int) {
	sum := 0
	for value := 0; value < num; value++ {
		if value%divider == 0 {
			sum += value
		}
	}
	resultChan <- sum
}

/**
创建简单的并行计算
*/
func test2() {
	fmt.Println("======创建简单的并行计算")
	go list_elem(10)

	//go运行完后就结束了，所以不会打印出list_elem元素，所以需要给它点时间
	//其实在test1中的sum3, sum5 := <-resultChan, <-resultChan，就是阻塞语句，当sum3,sum5取到值以后才会继续运行
	var input string
	fmt.Scanln(&input)
}

func list_elem(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

/**
为了看到并行的效果，改良一下test2
*/
func test3() {
	fmt.Println("======改良后的test2，验证是否真的并行了")
	go list_elem1(10, "go one")
	go list_elem1(20, "go two")

	var input string
	fmt.Scanln(&input)
}

func list_elem1(n int, name string) {
	for i := 0; i < n; i++ {
		fmt.Println(name, i)
		tick := time.Duration(rand.Intn(1000))
		time.Sleep(tick)
	}
}
