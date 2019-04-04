package server

import "errors"

type DemoService struct {
}

type Args struct {
	A, B int64
}

func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("can not division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
