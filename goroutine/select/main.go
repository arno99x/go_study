package main

import (
	"fmt"
	"time"
)

func createIntChan() chan int {
	c := make(chan int)
	i := 0
	go func() {
		for {
			time.Sleep(time.Millisecond * 800)
			c <- i
			i++
		}
	}()
	return c
}

func createWork() chan string {
	worker := make(chan string)
	return worker
}
func inputWork(n int, worker chan string, chanName string) {
	v := fmt.Sprintf("the value is %d from chan %s", n, chanName)
	worker <- v
}
func main() {
	c1, c2 := createIntChan(), createIntChan()
	worker := createWork()

	tm := time.After(time.Second * 10)
	tx := time.Tick(time.Second)
	for {
		select {
		//10秒后关闭应用
		case <-tm:
			fmt.Println("bye ...")
			return
		//schedule：每间隔两2秒打印一次
		case <-tx:
			fmt.Println("间隔2秒打印。。。")
		case n := <-c1:
			go inputWork(n, worker, "c1")
		case n := <-c2:
			go inputWork(n, worker, "c2")
		case v := <-worker:
			fmt.Println("out put from cache : ", v)
		//300毫秒没收到任何chan的值就打印这条
		case <-time.After(time.Millisecond * 300):
			fmt.Println("time out")
		}
	}
}
