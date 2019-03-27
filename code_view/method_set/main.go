package main

import "fmt"

type Do interface {
	Speak(content string)
}

type Cat struct {
	Name  string
	Color string
}

func (c Cat) Speak(content string) {
	fmt.Printf("%s speak : %s\n", c.Name, content)
}
func (d *Dog) Speak(content string) {
	fmt.Printf("%s speak : %s\n", d.Name, content)
}

type Dog struct {
	Name  string
	Color string
}

func doSpeak(do Do, content string) {
	do.Speak(content)
}

type BlackCat struct {
	Cat
	Age int
}

type WhiteCat struct {
	Cat
	Age int
}

func (wc WhiteCat) Speak(content string) {
	content = "this is white " + wc.Cat.Name + " " + content
	fmt.Println("inner : ", content)
}

func main() {
	c1 := Cat{Name: "c1", Color: "black"}
	doSpeak(c1, "喵喵喵~~~~~~")
	c2 := &Cat{Name: "c2", Color: "black"}
	doSpeak(c2, "喵喵喵~~~~~~")

	//说明Dog的speak()是绑定指针型Dog的，没要以传地址的方式调用，golang默认都是传值
	d1 := Dog{Name: "d1", Color: "black"}
	//doSpeak(d1)  //错误
	doSpeak(&d1, "汪汪汪~~~~~~") //正确
	d2 := &Dog{Name: "d2", Color: "black"}
	doSpeak(d2, "汪汪汪~~~~~~")

	//重点： black并未实现Do interface，这时是函数使用内部Cat的方法:func (c Cat)Speak(content string)
	bc := BlackCat{
		Cat: Cat{Name: "cat", Color: "black"},
		Age: 2,
	}
	doSpeak(bc, "black 汪汪汪~~~~~~~")
	doSpeak(bc.Cat, ".... 汪汪汪~~~~~~~~")

	//WhiteCat自实现一个Speak()方法看一下
	wc := WhiteCat{
		Cat: Cat{Name: "cat", Color: "white"},
		Age: 3,
	}
	//whiteCat的方法Speak覆盖了Cat的
	doSpeak(wc, "吼吼~~~")
	//也可以这样调用Cat的Speak方法
	doSpeak(wc.Cat, "吱吱吱~~~~~~")

}
