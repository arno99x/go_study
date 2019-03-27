package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second * 2)
			ch <- "result"
		}
	}()

	for {
		select {
		case res := <-ch:
			fmt.Println(res)
		case <-time.After(time.Second * 1):
			fmt.Println("timeout")
		}
	}

}
