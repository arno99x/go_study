package engine

import (
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"go_study/crawler_concurrence/fetcher"
)

type SimpleEngine struct {
}

func init() {
	fmt.Println("init ")
}
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		requestResult, err := Worker(request)
		if err != nil {
			continue
		}
		if requestResult.Requests != nil {
			requests = append(requests, requestResult.Requests...)
		}

		for _, item := range requestResult.Items {
			fmt.Printf("Got Item %v\n", item)

		}
	}
}

func Worker(request Request) (ParserResult, error) {
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Error(err)
		return ParserResult{}, err
	}
	parserResult := request.ParserFun(body)
	return parserResult, nil
}
