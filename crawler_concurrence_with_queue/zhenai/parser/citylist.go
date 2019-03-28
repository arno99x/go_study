package parser

import (
	"go_study/crawler_concurrence_with_queue/engine"
	"regexp"
)

const ParserCityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(content []byte) engine.ParserResult {
	r := regexp.MustCompile(ParserCityListRe)
	cityList := r.FindAllSubmatch(content, -1)

	result := engine.ParserResult{}
	//limit := 0
	for _, item := range cityList {
		//limit++
		//if limit > 3 {
		//	break
		//}
		result.Items = append(result.Items, string(item[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(item[1]), ParserFun: ParserCity})
	}

	return result

}
