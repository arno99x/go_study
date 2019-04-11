package main

import (
	"fmt"
	"os"
)

func main() {
	//打印不换行
	fmt.Print(".......")
	//打印换行
	fmt.Println("----------")
	//格式化输出
	i := 2
	s := "str"
	fmt.Printf("输出整型 %d", i)
	fmt.Printf("输出字符串 %s\n", s)
	fmt.Printf("如果你不想判断类型用 %v %v\n", i, s)
	fmt.Printf("")

	//转换
	f := 3.1415926
	iStr := fmt.Sprintf("%f", f)
	fmt.Println(iStr)

	//
	i, err := fmt.Fprintln(os.Stdout, "\n ppppppppppp")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("i = ", i)
	}

}
