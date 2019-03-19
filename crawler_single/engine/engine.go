package engine

import (
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"go_study/crawler_single/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Error(err)
			continue
		}
		parserResult := request.ParserFun(body)
		if parserResult.Requests != nil {
			requests = append(requests, parserResult.Requests...)
		}

		for _, item := range parserResult.Items {
			fmt.Printf("Got Item %v\n", item)

		}
	}
}
