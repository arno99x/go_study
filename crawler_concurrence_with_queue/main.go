package main

import (
	"fmt"
	"go_study/crawler_concurrence_with_queue/engine"
	"go_study/crawler_concurrence_with_queue/scheduler"
	"go_study/crawler_concurrence_with_queue/zhenai/parser"
)

func init() {
	fmt.Println("init main.go")
}
func main() {
	e := engine.ConcurrenceEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
	}

	e.Run(
		engine.Request{
			Url:       "http://www.zhenai.com/zhenghun",
			ParserFun: parser.ParserCityList,
		},
	)
}
