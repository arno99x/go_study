package main

import (
	"fmt"
	"regexp"
)

const text = `arno's email is arno.wang@gmail.com
				lisa's email is lisa.li@163.com
				lucas's email is lucas.liu@kii.com
				alfred's email is alfred@sina.com.cn
`

func main() {
	//正则表达式解析
	//[0-9a-zA-Z]   表示0-9或a-z或A-Z中的一个或多个字符
	//[0-9a-zA-Z]+  表示一个或多个

	//@[0-9a-zA-Z]+  表示一个@与0-9或a-z或A-Z中的一个或多个字符
	//\.             表示 . 因为.是正则保留字符，它表示任意字符，如果想用来单纯的表示字符.  ， 需要用转意符\
	r := regexp.MustCompile(`[0-9a-zA-Z]+@[0-9a-zA-Z]+\.[0-9a-zA-Z]+`)

	//匹配所有满足正则的文本
	matchs := r.FindAllString(text, -1)
	fmt.Println("match list 1 : ", matchs)

	//alfred@sina.com.cn这里输出有点问题，把.cn给漏掉了
	//分析，应该允许@后面重复出现   [0-9a-zA-Z.]+  并以 \.[0-9a-zA-Z]+ 结尾似乎就满足我们的需求了，现试一下吧
	r = regexp.MustCompile(`[0-9a-zA-Z]+@[0-9a-zA-Z.]+\.[0-9a-zA-Z]+`)
	matchs = r.FindAllString(text, -1)
	fmt.Println("match list 2 : ", matchs)

	//只想提取邮件的名称怎么办呢？
	//分两步： 1，把想提取的子串用（）给括起来 ； 2，用r.FindAllStringSubmatch(text,-1)来匹配
	r = regexp.MustCompile(`([0-9a-zA-Z]+)@([0-9a-zA-Z.]+\.[0-9a-zA-Z]+)`)
	subMatchs := r.FindAllStringSubmatch(text, -1)
	fmt.Println("subMatch list 2 : ", subMatchs)

	//匹配一个满足正则表达式的词
	match := r.FindString(text)
	fmt.Println(match)

}
