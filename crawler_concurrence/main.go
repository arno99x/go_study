package main

import (
	"fmt"
	"go_study/crawler_concurrence/engine"
	"go_study/crawler_concurrence/scheduler"
	"go_study/crawler_concurrence/zhenai/parser"
)

func init() {
	fmt.Println("init main.go")
}
func main() {
	e := engine.ConcurrenceEngine{
		Scheduler: &scheduler.SimpleSchedule{},
		WorkCount: 1,
	}

	e.Run(
		engine.Request{
			Url:       "http://www.zhenai.com/zhenghun",
			ParserFun: parser.ParserCityList,
		},
	)
}
