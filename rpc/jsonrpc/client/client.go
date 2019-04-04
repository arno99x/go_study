package main

import (
	"fmt"
	"go_study/rpc/jsonrpc/server"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var result float64
	client := jsonrpc.NewClient(conn)
	client.Call("DemoService.Div", server.Args{4, 0}, &result)
	fmt.Println(err, "   :   ", result)
}
