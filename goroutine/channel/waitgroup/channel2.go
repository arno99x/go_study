package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	in chan int
	wg *sync.WaitGroup
}

func (work Worker) DoWork(wg *sync.WaitGroup) {
	for c := range work.in {
		fmt.Printf("get data : %c\n", c)
		wg.Done()
	}
}

func CreateWorker(wg *sync.WaitGroup) (worker Worker) {
	worker = Worker{
		in: make(chan int),
		wg: wg,
	}
	go worker.DoWork(wg)
	return
}

func main() {
	var wg sync.WaitGroup
	wg.Add(20)

	work := CreateWorker(&wg)
	for i := 0; i < 10; i++ {
		work.in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		work.in <- 'A' + i
	}

	wg.Wait()
}
