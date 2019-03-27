package main

import (
	"fmt"
	"go_study/crawler_single/engine"
	"go_study/crawler_single/zhenai/parser"
)

func init() {
	fmt.Println("init main.go")
}
func main() {
	engine.Run(
		engine.Request{
			Url:       "http://www.zhenai.com/zhenghun",
			ParserFun: parser.ParserCityList,
		},
	)
}
