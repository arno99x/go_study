package main

import "fmt"

/**
工厂方法模式迷补了简单工厂模式的不足：
	新增工厂时不用修改工厂类了，做横向扩展即可，只要满足Annimal interface即可被扩展，符合开闭原则

不足之处： 当产品种类太多时会产生很多个产品工厂，这是很不好的
解决办法：可以用简单工厂与工厂方法结合的方式来处理，把产品树中的同类产品用简单工厂来处理

但是而对更复杂的情况，若存在不同的产品树，产品树上又存在不同的产品簇，这里用抽像工厂更合适一些

到底什么是产品簇？
举例说明：
	汽车产品面有奥迪、宝马
	奥迪这个产品树下有运动车型和商务车型
	宝马这个产品树下也有运动车型和商务车型

这种情况下可以把所有运动车型归为一个产品簇，所有商务车型归为一个产品簇
*/

type Annimal interface {
	Eat()
}

type Dog struct {
}

type Cat struct {
}

func (cat Cat) Eat() {
	fmt.Println("cat eat ... ")
}

func (dog Dog) Eat() {
	fmt.Println("dog eat ... ")
}

type DogFactory struct {
}

type CatFactory struct {
}

func (DogFactory) Create() Annimal {
	dog := Dog{}
	return dog
}

func (CatFactory) Create() Annimal {
	cat := Cat{}
	return cat
}

func main() {
	df := DogFactory{}
	df.Create().Eat()

	cf := CatFactory{}
	cf.Create().Eat()
}
