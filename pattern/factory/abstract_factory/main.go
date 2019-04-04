/**
抽象工厂模式：
	一般用于具有产品树和产品族的场景下。
抽象工厂模式的缺点：
	如果需要增加新的产品树Audi，那么我们需要怎么做呢？ 你得先添加AudiCar AudiSportCar AudiBusinessCar,然后再去修改一下工厂方法。
	这么麻烦的事情也不是我们希望看到的

解决方法：可以用简单工厂配合反射来改进抽象工厂
	这个用JAVA比较好实现，GO目前不支持字过字符串反射来创建结构体，待后续探索，办法一定会有

*/

package main

import "fmt"

type Car interface {
	Driver()
}

type BmwSportCar struct {
}

func (BmwSportCar) Driver() {
	fmt.Println("BmwSportCar driver ...")
}

type BenzSportCar struct {
}

func (BenzSportCar) Driver() {
	fmt.Println("BenzSportCar driver ...")
}

type BmwBusinessCar struct {
}

func (BmwBusinessCar) Driver() {
	fmt.Println("BmwBusinessCar driver ...")
}

type BenzBusinessCar struct {
}

func (BenzBusinessCar) Driver() {
	fmt.Println("BenzBusinessCar driver ...")
}

type BusinessDriver struct {
}

func (bd BusinessDriver) CreateBmwCar() Car {
	return BmwBusinessCar{}
}

func (bd BusinessDriver) CreateBenzCar() Car {
	return BenzBusinessCar{}
}

type SportDriver struct {
}

func (bd SportDriver) CreateSportBmwCar() Car {
	return BmwSportCar{}
}

func (bd SportDriver) CreateSportBenzCar() Car {
	return BenzSportCar{}
}

func main() {
	SportDriver{}.CreateSportBenzCar().Driver()
	SportDriver{}.CreateSportBmwCar().Driver()

	BusinessDriver{}.CreateBenzCar().Driver()
	BusinessDriver{}.CreateBmwCar().Driver()
}
