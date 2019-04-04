package main

import (
	"fmt"
	"go_study/pattern/factory/abstract_factory/case1"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	uData := factory.User{1, "u"}
	dData := factory.Department{1, "d"}
	data := factory.DataAccess{}

	iU := data.CreateUser("access")
	iU.Insert(&uData)
	gU := iU.GetUser(1)
	fmt.Println(gU)

	fmt.Println("================")

	iD := data.CreateDepartment("sqlserver")
	iD.Insert(&dData)
	gD := iD.GetDepartment(1)
	fmt.Println(gD)

	if iS := data.CreateDepartment("a"); iS != nil {
		t.Error()
	}
}
