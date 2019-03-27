package main

import (
	"fmt"
	"time"
)

/**
演示GO的协程
得出结论：
1、 go的goroutine函数不需要预定义，需要的时候在前面加上go关键字就可以
2、main函数结束后，所有协程都会结束，所以要让主函数休息一会

*/
func main() {
	for i := 0; i < 10; i++ {
		go worker1(i)
	}

	var arr = [10]int{}
	for i := 0; i < 10; i++ {
		go func(i int, arr [10]int) {
			arr[i]++
		}(i, arr)
	}
	fmt.Printf("%v\n", arr)

	//var arr = [10]int{}
	//for i := 0;i<10 ;i++  {
	//	go worker2(i,&arr)
	//}
	//fmt.Printf("%v\n",arr)

	time.Sleep(time.Millisecond)
}

func worker1(i int) {
	fmt.Printf("this is %d\n", i)
}

func worker2(i int, arr *[10]int) {
	arr[i]++
}
