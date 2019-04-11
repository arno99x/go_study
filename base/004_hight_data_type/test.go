package main

import (
	"fmt"
)

func main() {
	//数组
	fmt.Println("-----array test------------------------------")
	array_test()
	//切片(Slice)
	fmt.Println("------slice test-----------------------------")
	slice_test()
	fmt.Println("------map test-----------------------------")
	//字典(Map)
	map_test()
}

/**
字典是一组无序的，键值对的集合。
*/
func map_test() {
	map_test1()
	map_test2()
	map_test3()
}

func map_test1() {
	//字典的定义也有两种，一种是初始化数据的定义方式，另一种是使用神奇的make函数来定义。
	fmt.Println("=====招式1：初始化数据")
	x := map[string]string{
		"name": "zhangsan",
		"sex":  "男",
		"age":  "9",
	}

	for k, v := range x {
		fmt.Println(k, v)
	}

	fmt.Println("=====招式2：make函数定义")
	y := make(map[string]string)
	y["name"] = "lishi"
	y["sex"] = "女"
	y["age"] = "8"

	for k, v := range y {
		fmt.Println(k, v)
	}

	fmt.Println("=====招式3：如果map对应的KEY不存在，则返回默认值， string对应\"\",int对应0")
	z := map[string]int{
		"A": 0,
		"B": 13,
		"C": 14}
	fmt.Println("z['D']", z["D"])

	fmt.Println("=====招式4：招式3的补充，招式3中的z[\"D\"与z[\"A\"两种情况就分不清了")

	e := map[string]int{
		"A": 0,
		"B": 13,
		"C": 14}

	if val, ok := e["A"]; ok {
		fmt.Println(val, "true")
	} else {
		fmt.Println(val, "false")
	}
}

func map_test2() {
	//Go提供的内置函数delete，这个函数可以用来从字典中删除元素。
	fmt.Println("=====招式5：删除字典元素")
	e := map[string]int{
		"A": 0,
		"B": 13,
		"C": 14}

	delete(e, "A")
	fmt.Println(e)

}

func map_test3() {
	//实例场景：我们有一个学生登记表，登记表里面有一组学号，每个学号对应一个学生，每个学生有名字和年龄。
	fmt.Println("=====实例场景========")
	x := make(map[string]map[string]int)

	x["1000001"] = map[string]int{"age": 21, "sex": 1}
	x["1000002"] = map[string]int{"age": 21, "sex": 0}
	x["1000003"] = map[string]int{"age": 21, "sex": 1}

	fmt.Println(x)

	for k, v := range x {
		fmt.Println(k, ":")
		for k1, v1 := range v {
			fmt.Println("	", k1, v1)
		}
	}

	//也可以直接赋值
	y := map[string]map[string]int{
		"100001": {"age": 23, "sex": 1},
		"100002": {"age": 22, "sex": 0},
		"100003": {"age": 21, "sex": 1},
	}
	fmt.Println(y)

	for k, v := range y {
		fmt.Println(k, ":")
		for k1, v1 := range v {
			fmt.Println("	", k1, v1)
		}
	}
}

/**
切片和数组很类似，甚至你可以理解成数组的子集。但是切片有一个数组所没有的特点，那就是切片的长度是可变的。
严格地讲，切片有容量(capacity)和长度(length)两个属性。
*/
func slice_test() {
	slice_test1()
	slice_test2()
	slice_test3()
}

func slice_test1() {
	//首先我们来看一下切片的定义。切片有两种定义方式，一种是先声明一个变量是切片，然后使用内置函数make去初始化这个切片。另外一种是通过取数组切片来赋值。
	fmt.Println("=====招式1：声明一个变量是切片")
	x := make([]float64, 5)
	fmt.Println("capcity:", cap(x), "length:", len(x))

	y := make([]float64, 5, 10)
	fmt.Println("capcity:", cap(y), "length:", len(y))

	for i := 0; i < len(x); i++ {
		x[i] = float64(i * i)
	}
	fmt.Println(x)

	//虽然切片的容量可以大于长度，但是赋值的时候要注意最大的索引仍然是len(x)－1。否则会报索引超出边界错误。
	for i := 0; i < len(y); i++ {
		y[i] = float64(i + 6)
	}
	fmt.Println(y)

	fmt.Println("=====招式2：通过取数据切片来赋值")
	arr := [5]int{1, 2, 3, 4, 5}

	s1 := arr[2:4]
	fmt.Println(s1)

	s2 := arr[3:]
	fmt.Println(s2)

	s3 := arr[:3]
	fmt.Println(s3)

	s4 := arr[:]
	fmt.Println(s4)
}

func slice_test2() {
	//切片的长度可以变化
	fmt.Println("=====招式3：切处长度是可变的")
	arr := make([]int, 5, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 1
	}
	fmt.Println(arr)

	//Go在默认的情况下，如果追加的元素超过了容量大小，Go会自动地重新为切片分配容量，容量大小为原来的两倍
	arr = append(arr, 10, 11, 12, 13, 14, 15, 16)
	fmt.Println("Capacity:", cap(arr), "Length:", len(arr))
	fmt.Println(arr)
}

func slice_test3() {
	//copy函数用来从一个切片拷贝元素到另一个切片
	fmt.Println("=====招式3：从一个切片拷贝元素到另一个切片")
	s1 := []float64{10, 11, 12, 13, 14, 15}

	s2 := make([]float64, 5, 10)

	//fmt.Println(s1,s2)

	copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)

}

/**
数组是一个具有相同数据类型的元素组成的固定长度的有序集合
*/
func array_test() {
	array_test1()
	array_test2()
	array_test3()
	array_test4()
	array_test5()
}

func array_test1() {
	fmt.Println("=====招式1：分离赋值")
	//表示数组x是一个整型数组，而且数值的长度为5。
	var x [5]int
	x[0] = 2
	x[1] = 3
	x[2] = 3
	x[3] = 2
	x[4] = 12

	var sum int
	for _, elem := range x {
		sum += elem
	}
	fmt.Println(sum)
}

func array_test2() {
	fmt.Println("=====招式2：可以分开赋值")
	var x = [5]int{1, 2, 3, 4}
	x[4] = 5

	var sum int
	for _, elem := range x {
		sum += elem
	}
	fmt.Println(sum)
}

func array_test3() {
	fmt.Println("=====招式3：")
	var x = [5]int{}
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	var sum int
	for _, i := range x {
		sum += i
	}
	fmt.Println(sum)
}

func array_test4() {
	fmt.Println("招式4：不显示指定数组长度")
	var x = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday"}
	for _, day := range x {
		fmt.Println(day)
	}
}
func array_test5() {
	fmt.Println("招式5：对比招式4，最后一行后面后面若不跟上}号就要加个,号")
	var x = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	for _, day := range x {
		fmt.Println(day)
	}
}
