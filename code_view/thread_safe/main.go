package main

import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	ua := UserAges{ages: map[string]int{"zhangshan": 20, "lisi": 21, "wangwu": 22}}
	go func() {
		ua.Add("wangerma", 19)
		ua.Add("maliu", 18)
		ua.Add("qinyu", 17)
	}()

	go func() {
		fmt.Printf("%d", ua.Get("qinyu"))
		fmt.Printf("%d", ua.Get("maliu"))
		fmt.Printf("%d", ua.Get("wangerma"))
		fmt.Printf("%d", ua.Get("wangwu"))
		fmt.Printf("%d", ua.Get("lisi"))
		fmt.Printf("%d", ua.Get("zhangshan"))
	}()

	time.Sleep(time.Second)
}
