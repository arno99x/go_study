package main

import "fmt"

func main() {
	result := add(1, 2, 3)
	fmt.Println(result)

	result = add([]int{1, 2, 3}...)
	fmt.Println(result)

	pase_student()

}

func add(args ...int) int {
	sum := 0
	for _, s := range args {
		sum += s
	}
	return sum
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Printf("k = %s ; v = %d\n", k, v)
	}

}
