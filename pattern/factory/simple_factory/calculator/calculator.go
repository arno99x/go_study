/**
利用简单工厂模式来完成计算器的程序
*/
package calculator

import (
	"errors"
)

type Operator interface {
	Operator(a int, b int) (int, error)
}

type PlusCalulator struct {
}

type MinusCalulator struct {
}

type MultiplyCalulator struct {
}

type DevideCalulator struct {
}

func (PlusCalulator) Operator(a int, b int) (int, error) {
	return a + b, nil
}

func (MinusCalulator) Operator(a int, b int) (int, error) {
	return a - b, nil
}

func (MultiplyCalulator) Operator(a int, b int) (int, error) {
	return a * b, nil
}

func (DevideCalulator) Operator(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("can not divide by 0")
		return 0, err
	}
	return a / b, nil
}

type Factory interface {
	GetOperator(operater string)
}

func GetOperator(operator string) Operator {
	switch operator {
	case "+":
		return PlusCalulator{}
	case "-":
		return MinusCalulator{}
	case "*":
		return MultiplyCalulator{}
	case "/":
		return DevideCalulator{}
	default:
		return nil

	}
}

func ToCalulator(a int, b int, operator string) (int, error) {
	calulator := GetOperator(operator)
	if calulator == nil {
		return 0, errors.New("can not support this operator")
	}
	return calulator.Operator(a, b)
}
