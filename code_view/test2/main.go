package main

import (
	"fmt"
	"runtime"
)

/**
select随机性:
select会随机选择一个可用通用做收发操作。 所以代码是有肯触发异常，也有可能不会。 单个chan如果无缓冲时，将会阻塞。

但结合 select可以在多个chan间等待执行。有三点原则：
1、select 中只要有一个case能return，则立刻执行。
2、当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
3、如果没有一个case能return则可以执行”default”块。
*/
func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
