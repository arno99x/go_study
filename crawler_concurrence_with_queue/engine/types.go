package engine

type Request struct {
	Url       string
	ParserFun func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
