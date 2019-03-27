package main

import (
	"fmt"
	"sync"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}) // 解除注释看看！
	//ch := make(chan interface{},len(set.s))
	go func() {
		//set.RLock()

		for elem, value := range set.s {
			println("-> Iter:", elem, value)
			ch <- elem

		}

		close(ch)
		//set.RUnlock()

	}()
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	vc := th.Iter()
	for c := range vc {
		fmt.Printf(" ===== %d\n", c)
	}

	//fmt.Sprintf("%s ------ %v","ch",vc)
}
