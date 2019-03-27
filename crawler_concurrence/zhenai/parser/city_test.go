package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
		return
	}
	result := ParserCity(contents)

	for index, item := range result.Items {
		fmt.Printf("\n%s\n", item)
		fmt.Printf(result.Requests[index].Url)
	}
}
