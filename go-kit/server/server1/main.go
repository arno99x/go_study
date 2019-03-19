package main

import (
	"fmt"
	"go_study/go-kit/server/server1/service"
	"go_study/go-kit/server/server1/service/dbclient"
	"go_study/go-kit/server/server1/service/handle"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient() // NEW
	service.StartWebServer("6767")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	handle.DBClient = &dbclient.BoltClient{}
	handle.DBClient.OpenBoltDb()
	handle.DBClient.Seed()
}
