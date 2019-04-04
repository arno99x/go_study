package rpcdemo

import (
	"go_study/rpc/jsonrpc/server"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
启动后的测试入参：
telnet local host 1234
回车
{"method":"DemoService.Div","params":[{"A":3,"B":0}],"id":1}
*/
func main() {
	rpc.Register(server.DemoService{})
	listen, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("err info : %s", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
