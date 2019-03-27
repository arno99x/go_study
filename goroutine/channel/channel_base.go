package main

import (
	"fmt"
	"time"
)

var ch chan int

func helloChan() {
	fmt.Println("hello chan !")

	ch <- 1
	ch <- 2

	time.Sleep(time.Second * 5)
	ch <- 3
}

func main() {
	//通道申明
	ch = make(chan int)

	go helloChan()

	value := <-ch
	fmt.Printf("csdccc %d\n", value)

	value = <-ch
	fmt.Printf("csdccc %d\n", value)

	//go helloChan()
	fmt.Println("==============")
	value = <-ch
	fmt.Printf("tttt %d\n", value)

	//value = <- ch
	//fmt.Printf("tttt %d\n",value)

	//time.Sleep(250*time.Millisecond)
}
