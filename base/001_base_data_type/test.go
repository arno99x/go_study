package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	_data()
	_string()
	_bool()
}

func init() {
	fmt.Println("init test.go")
}
func _bool() {
	fmt.Println("------bool test----------------------")
	var equal bool
	var a int = 10
	var b int = 20
	equal = (a == b)
	fmt.Println(equal)
}

func _string() {
	str_test1()
	str_test2()
}

/**
字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。（对于汉字，通常由多个字节组成）。
这就是说，传统的字符串是由字符组成的，而Go的字符串不同，是由字节组成的。这一点需要注意。

字符串的表示很简单。用(双引号””)或者(``号)来描述。

唯一的区别是，双引号之间的转义字符会被转义，而``号之间的转义字符保持原样不变。
*/
func str_test1() {
	fmt.Println("------str_test1---------------------------")
	fmt.Println("GO字符串可以由单引号或双引号来表示，唯一区别就是单引号会保持字符原样")
	var a string = "hello \n world \n"
	var b string = `hello \n world \n`
	fmt.Println(a)
	fmt.Println(b)
}

/**
  字符串所能进行的一些基本操作
*/
func str_test2() {
	var a string = "1234567890"
	var b string = "-"
	var c string = "hahahaha"

	fmt.Println("------str_test2---------------------------")
	fmt.Println("=====功能1：获取字符串长度")
	fmt.Println(len(a))

	fmt.Println("=====功能2：获取字符串中单个字节")
	fmt.Println("这里我们看到a[9]得到的是一个整数，这就证明了上面\"Go的字符串是由字节组成的这句话\"。")
	fmt.Println(a[9])

	fmt.Println("=====功能3：字符串联接")
	fmt.Println(a + b + c)

	fmt.Println("功能4：int转换为string")
	var d uint8 = a[9]
	fmt.Println(string(d))

	fmt.Println("=====功能5：字符串比较")
	var str01 string = "abc"
	var str02 string = "abc"
	com01 := strings.Compare(str01, str02)
	if com01 == 0 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等 " + string(com01))
	}
	fmt.Println(com01)

	fmt.Println("=====功能6：查找 包含")
	str01 = "helloworld"
	var isCon bool = strings.Contains(str01, "hello")
	fmt.Println(isCon) //true

	fmt.Println("=====功能7：查找位置")
	//查找位置
	var theIndex int = strings.Index(str01, ",")
	fmt.Println(theIndex)                     //5
	fmt.Println(strings.Index(str01, "haha")) //不存在返回-1

	lastIndex := strings.LastIndex(str01, "o")
	fmt.Println("在字符串中最后出现位置的索引 " + strconv.Itoa(lastIndex)) //7
	//-1 表示字符串 s 不包含字符串

	fmt.Println("=====功能8：字符分隔")
	//根据特定字符分割
	slice01 := strings.Split("q,w,e,r,t,y,", ",")
	fmt.Println(slice01)      //[q w e r t y ]
	fmt.Println(cap(slice01)) //7  最后多个空""
	for i, v := range slice01 {
		fmt.Printf("下标 %d 对应值 = %s \n", i, v)
	}

	fmt.Println("字符串分割")
	fieldsStr := "  hello   it's  a  nice day today    "
	//根据空白符分割,不限定中间间隔几个空白符
	fieldsSlece := strings.Fields(fieldsStr)
	fmt.Println(fieldsSlece) //[hello it's a nice day today]

	for i, v := range fieldsSlece {
		fmt.Printf("下标 %d 对应值 = %s \n", i, v)
	}
	for i := 0; i < len(fieldsSlece); i++ {
		fmt.Println(fieldsSlece[i])
	}

	fmt.Println("=====功能9：替换")
	// 在s字符串中, 把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	var str03 string = "/Users//Documents/GOPatch/src/MyGO/config/TestString/"
	str04 := strings.Replace(str03, "/", "**", -1)
	str05 := strings.Replace(str03, "/", "**", 4)

	fmt.Println(str04) //**Users****Documents**GOPatch**src**MyGO**config**TestString**
	fmt.Println(str05) //**Users****Documents**GOPatch/src/MyGO/config/TestString/

	fmt.Println("=====功能10：前缀 后缀")
	//前缀 后缀
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasSuffix("Amigo", "go"))  // true

	fmt.Println("功能11：字符串拼接")
	//Join 用于将元素类型为 string 的 slice, 使用分割符号来拼接组成一个字符串：
	var str08 string = strings.Join(fieldsSlece, ",")
	fmt.Println("Join拼接结果=" + str08) //hello,it's,a,nice,day,today
}

func _data() {
	float_test()
	int_test()
	uint_test()
	uintptr_test()
	complex_test()
}

/**
Go的浮点数类型有两种，float32和float64。float32又叫单精度浮点型，float64又叫做双精度浮点型。其最主要的区别就是小数点后面能跟的小数位数不同。
*/
func float_test() {
	fmt.Println("------float_test---------------------------")
	var a float32 = 12.34
	var b float32 = 3.45
	fmt.Println(a / b)
	//支持合更转换
	fmt.Println(int(a))
}

/**
 * GO中三个依懒系统的数据类型 int , uint , uintptr
 * 它们在32、64位操作系统中表示的范围不一样，32位中表示的是int32、uint32、uintptr32
 */
func int_test() {
	fmt.Println("------int_test-----------------------------")
	fmt.Println("int类型还有int8/int16/int32/int64, 位数越大说明能表示的数值范围越大")
	var a int = 6
	var b int = 7
	fmt.Println(a + b)
}
func uint_test() {
	fmt.Println("------uint_test-----------------------------")
	fmt.Println("uint类型还有uint8/uint16/uint32/uint64, 位数越大说明能表示的数值范围越大")
	var a uint = 9
	var b uint = 7
	fmt.Println(a + b)
}
func uintptr_test() {
	fmt.Println("------uintptr_test-----------------------------")
	fmt.Println("uintptr类型没有指定位数表示")
	var a uintptr = 9
	var b uintptr = 7
	fmt.Println(a + b)
}

/**
  GO语言独有的虚数类型，有complex64、complex128两种
*/
func complex_test() {
	fmt.Println("------complex_test-----------------------------")
	var a complex64 = 64
	var b complex64 = 64
	fmt.Println(a + b)
}
