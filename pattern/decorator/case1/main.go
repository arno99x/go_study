/**
装饰者模式：动态的给一些对象添加一些职责，就增加功能来说，装饰模式比增加子类更加灵活


*/
package main

import "fmt"

type Component interface {
	Operation()
}

type Decorator interface {
	Operation()
	SetComponent()
}

type ConcreteComponent struct{}

func (concreteComponent ConcreteComponent) Operation() {
	fmt.Println("对象的操作 。。。 ")
}

type ConcreteComponentA struct {
	Component   Component
	DoOperation func()
}

func (concreteComponentA *ConcreteComponentA) SetComponent(component Component) {
	concreteComponentA.Component = component
}
func (concreteComponentA ConcreteComponentA) Operation() {
	concreteComponentA.Component.Operation()
	fmt.Println("对象A的操作 。。。")
}

type ConcreteComponentB struct {
	Component Component
}

func (concreteComponentB *ConcreteComponentB) SetComponent(component Component) {
	concreteComponentB.Component = component
}
func (concreteComponentB ConcreteComponentB) Operation() {
	concreteComponentB.Component.Operation()
	fmt.Println("对象B的操作 。。。")
}

func main() {
	concreteComponent := ConcreteComponent{}
	concreteComponentA := ConcreteComponentA{}
	concreteComponentB := ConcreteComponentB{}

	concreteComponentA.SetComponent(&concreteComponent)
	concreteComponentB.SetComponent(&concreteComponentA)

	concreteComponentB.Operation()
}
