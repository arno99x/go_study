package yellow

import (
	"strconv"
)

type Duck struct {
	Color       string
	Weight      float64
	YellowValue string
}

func (duck Duck) GetInfo() string {
	return duck.Color + " , " + strconv.FormatFloat(duck.Weight, 'E', -1, 64)
}
