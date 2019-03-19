package main

import "fmt"

//详细说明参见我的博客：https://blog.csdn.net/wangqiang9x/article/details/86222313
func main() {
	s := []int{0, 1, 2, 3}
	fmt.Println("copy slice")
	s1 := make([]int, 8)
	copy(s1, s)
	printfSlice(s1)

	fmt.Println("delete slice middle element")
	s2 := append(s1[:3], s1[4:]...)
	fmt.Println(s2)

	fmt.Println("delete slice left element")
	s3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[1:]
	printfSlice(s4)

	fmt.Println("delete slice right element")
	s5 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s6 := s5[:len(s5)-1]
	printfSlice(s6)

}

func printfSlice(s []int) {
	fmt.Println("s=%v , len(s)=%d , cap(s)=%d", s, len(s), cap(s))
}
