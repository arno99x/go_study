package main

import (
	"fmt"
	"time"
)

type Worker struct {
	in   chan int
	done chan bool
}

func (work Worker) DoWork() {
	for c := range work.in {
		fmt.Printf("get data : %c\n", c)
		go func() {

			fmt.Println("insert to done channel ")
			work.done <- true
		}()
	}
}

func CreateWorker() (worker Worker) {
	worker = Worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go worker.DoWork()
	return
}

func main() {
	work := CreateWorker()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		work.in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		work.in <- 'A' + i
	}

	fmt.Println("read from done channel ", <-work.done)

	fmt.Println("============end================")
}
