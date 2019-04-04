/**
单例工厂模式
*/

package main

import (
	"fmt"
	"sync"
)

//饿汉模式
var sc1 *singletonCase1 = new(singletonCase1)

type singletonCase1 struct {
}

func (singletonCase1) show() {
	fmt.Println("this is singletonCase1 show method ... ")
}
func GetSingletonCase1() *singletonCase1 {
	return sc1
}

//懒汉模式
var sc2 *singletonCase2

type singletonCase2 struct{}

func (singletonCase2) show() {
	fmt.Println("this is singletonCase2 show method ... ")
}

//这种写法是存在问题的，有可能会存在两个调用同时进入if方法体，这样就给外面提供了两个实例
func GetSingletonCase20() *singletonCase2 {
	if sc2 == nil {
		sc2 = &singletonCase2{}
	}
	return sc2
}

//解决GetSingletonCase20的问题可以用锁
//虽然避免了问题，但每次加锁也大大降低了效率
var mu sync.Mutex

func GetSingletonCase21() *singletonCase2 {
	mu.Lock()
	defer mu.Unlock()
	if sc2 == nil {
		sc2 = &singletonCase2{}
	}
	return sc2
}

//解决GetSingletonCase21的缺陷可以用双重锁
//下面的代码看似重复，实则巧炒，只有在程序初始化的时候才可能会运行到外层的if代码段中，初始化完成后就不会被加锁了
func GetSingletonCase22() *singletonCase2 {
	if sc2 == nil {
		mu.Lock()
		defer mu.Unlock()
		if sc2 == nil {
			sc2 = &singletonCase2{}
		}
	}
	return sc2
}

//用sync.once() 加锁
var once sync.Once

func GetSingletonCase23() *singletonCase2 {
	once.Do(func() {
		sc2 = &singletonCase2{}
	})
	return sc2
}

func main() {
	GetSingletonCase1().show()

	GetSingletonCase20().show()
	GetSingletonCase21().show()
	GetSingletonCase22().show()
	GetSingletonCase23().show()
}
