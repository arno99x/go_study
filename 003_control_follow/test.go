package main

import (
	"fmt"
	"strconv"
)

func main() {
	if_else_test()
	switch_test()
	for_test()
}

func for_test() {
	fmt.Println("第一式：标准式")
	for i := 1; i < 4; i++ {
		fmt.Println("第" + strconv.Itoa(i) + "个")
	}

	fmt.Println("第二式：模坊while   <go里面没有while循环>")
	i := 1
	for i < 4 {
		fmt.Println("第" + strconv.Itoa(i) + "个")
		i++
	}

	fmt.Println("第三式：死循环")
	for {
		fmt.Println("死了。。。。")
	}
}

func switch_test() {
	fmt.Println("------switch test-----------------------")
	score := 68

	switch score / 10 {
	case 10:
		fmt.Println("10")
	case 9:
		fmt.Println("9")
	case 8:
		fmt.Println("8")
	case 7:
		fmt.Println("7")
	case 6:
		fmt.Println("6")
	case 5:
		fmt.Println("5")
	case 4:
		fmt.Println("4")
	case 3:
		fmt.Println("3")
	case 2:
		fmt.Println("2")
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("error")

	}
}
func if_else_test() {
	fmt.Println("------if else test-----------------------")
	dog_age := 7
	fmt.Println(strconv.Itoa(dog_age))
	if dog_age < 3 {
		fmt.Println("it a young dog , age = " + strconv.Itoa(dog_age))
	} else if dog_age >= 3 && dog_age <= 5 {
		fmt.Println("it a middle dog , age = " + strconv.Itoa(dog_age))
	} else {
		fmt.Println("it a old dog , age = " + strconv.Itoa(dog_age))
	}
}
