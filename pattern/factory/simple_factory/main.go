/**
简单工石模式：向外界屏蔽了直接创建的功能，由工厂来负责这个工作
好处：当增加一个产品时只需要在工厂类添加判断即可，对用户来说只要传入新的参数即可获得
不足之处：
1. 只有一个工厂类，但很多复杂的情况下产品结构是呈树树状结构的
2. 当加入新产品时需要修改核心的工厂类，这存在风险，也不满足开闭原则

根据上面的不足，引入工厂方法模式

注：
开闭原则 ： 对扩展开放，对修改关闭

*/
package main

import "fmt"

type Annimal interface {
	Eat()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) Eat() {
	fmt.Println("dog eat ... ")
}

func (cat Cat) Eat() {
	fmt.Println("cat eat ... ")
}

func CreateFactory(typeName string) Annimal {
	switch typeName {
	case "dog":
		return Dog{}
	case "cat":
		return Cat{}
	default:
		return nil
	}
}

func main() {
	annimal := CreateFactory("dog")
	if annimal == nil {
		fmt.Println("can not find any annimal")
	} else {
		annimal.Eat()
	}

	annimal = CreateFactory("cat")
	if annimal == nil {
		fmt.Println("can not find any annimal")
	} else {
		annimal.Eat()
	}

	annimal = CreateFactory("mouse")
	if annimal == nil {
		fmt.Println("can not find any annimal")
	} else {
		annimal.Eat()
	}

}
