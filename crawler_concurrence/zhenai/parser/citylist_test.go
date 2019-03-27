package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParserCityList(contents)
	//fmt.Printf("\n length -> %s\n",strconv.Itoa(len(result.Requests)))
	//for len(result.Requests)>0 {
	//
	//	fmt.Println(result.Requests[0])
	//	result.Requests = result.Requests[1:]
	//}
	const requestSize = 471
	if len(result.Requests) != requestSize {
		t.Errorf("this city size is %d not incorrent !", len(result.Requests))
	}
}
