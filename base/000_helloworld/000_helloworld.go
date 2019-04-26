//main包, 凡是标注为main包的go文件都会被编译为可执行文件
package main

//导入需要使用的包
import (
	"fmt" //支持格式化输出的包,就是format的简写
	"os"
)

//主函数,程序执行入口
func main() {
	/*
	   输出一行hello world
	   Println函数就是print line的意思
	*/

	fmt.Println("hello world")
	fmt.Println("*** " + os.Getenv("PATH"))
}
