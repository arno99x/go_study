/**
抽像工厂模式再发一下工厂方法模式的差别是，抽像工厂更拔高了一层，可以把一些工厂方法给聚类，变成一个系列的产品

可以看各种产品与各种工厂是松散的分开的，通过对工厂的聚合来达到产品分类的目的

通过工厂模式写出来的代码是耦合度低的，易扩展的，健壮的，可维护的
*/
package factory

import (
	"fmt"
)

type Food interface {
	Eat()
}

type Drink interface {
	Drink()
}

type Hamberger struct {
	info string
}

type Meat struct {
	info string
}

type Coco struct {
	info string
}

type Cafe struct {
	info string
}

func (Hamberger) Eat() {
	fmt.Println("eat hamberger ...")
}

func (Meat) Eat() {
	fmt.Println("eat meat ...")
}

func (Coco) Drink() {
	fmt.Println("drink coco ...")
}

func (Cafe) Drink() {
	fmt.Println("drink cafe ...")
}

type FastFoodFactory interface {
	CreateFood()
	CreateDrink()
}

type FastFoodFactoryA struct {
}
type FastFoodFactoryB struct {
}

func (FastFoodFactoryA) CreateFood() Food {
	return Meat{"五分熟"}
}

func (FastFoodFactoryA) CreateDrink() Drink {
	return Coco{}
}

func (FastFoodFactoryB) CreateFood() Food {
	return Hamberger{}
}

func (FastFoodFactoryB) CreateDrink() Drink {
	return Cafe{}
}
