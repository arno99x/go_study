package main

import (
	"go_study/crawler_single/engine"
	"go_study/crawler_single/zhenai/parser"
)

func main() {
	engine.Run(
		engine.Request{
			Url:       "http://www.zhenai.com/zhenghun",
			ParserFun: parser.ParserCityList,
		},
	)
}
