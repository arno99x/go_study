package parser

import (
	"go_study/crawler_single/engine"
	"go_study/crawler_single/model"
	"regexp"
	"strconv"
	"strings"
)

const purpleRe = `<div class="m-btn purple" [^>]*>([^<]+)</div>`
const nameRe = `<h1 class="nickName"[^>]*>([^<]+)</h1>`
const baseInfoRe = `<div class="des f-cl"[^>]*>([^>]+)</div>`

func ParserProfile(contents []byte, paramMap map[string]string) engine.ParserResult {
	profile := model.Profile{}

	nameRege := regexp.MustCompile(nameRe)
	nameMatch := nameRege.FindAllSubmatch(contents, -1)
	//fmt.Printf("\nname-> %s\n",nameMatch[0][1])
	profile.Name = string(nameMatch[0][1])
	profile.Gender = paramMap["gender"]
	baseInfoRege := regexp.MustCompile(baseInfoRe)
	baseInfoMatch := baseInfoRege.FindAllSubmatch(contents, -1)
	//fmt.Printf("\nbaseInfo -> %s\n",baseInfoMatch[0][1])
	baseInfoList := strings.Split(string(baseInfoMatch[0][1]), "|")
	//fmt.Println(baseInfoList)
	profile.HuKou = baseInfoList[0]
	profile.Age, _ = strconv.Atoi(baseInfoList[1])
	profile.Education = baseInfoList[2]
	profile.Marriage = baseInfoList[3]
	profile.Height, _ = strconv.Atoi(baseInfoList[4])
	profile.Income = baseInfoList[5]

	purleRegexp := regexp.MustCompile(purpleRe)
	matchs := purleRegexp.FindAllSubmatch(contents, -1)
	for len(matchs) > 0 {
		//fmt.Printf("%s\n",matchs[0][1])
		//获取星座
		xinzuo := parserXinzuo(matchs[0])
		if xinzuo != "" {
			profile.XinZuo = xinzuo
		}
		//获取体重
		weight := parserWeight(matchs[0])
		if weight != "" {
			weightValue, err := strconv.Atoi(weight)
			if err != nil {
				profile.Weight = weightValue
			}
		}

		matchs = matchs[1:]
	}

	result := engine.ParserResult{Items: []interface{}{profile}, Requests: nil}
	return result
}

//获取星座
func parserXinzuo(matchs [][]byte) string {
	for e := range model.XinZus {
		if strings.Contains(string(matchs[1]), model.XinZus[e]) {
			return model.XinZus[e]
		}
	}
	return ""
}

//获取体重
const weightRe = `[0-9]+kg`

func parserWeight(matchs [][]byte) string {
	//weightRegexp := regexp.MustCompile(weightRe)
	//match := weightRegexp.FindSubmatch(matchs[1])
	//fmt.Printf("%s",match[0])
	return ""
}
