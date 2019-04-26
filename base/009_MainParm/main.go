package main

import (
	"fmt"
	"github.com/wonderivan/logger"
	"os"
)

//import (
//	"flag"
//)
//
//var strFlag = flag.String("s", "", "Description")
//var boolFlag = flag.Bool("bool", false, "Description of flag")
//
//func main() {
//	flag.Parse()
//	println(*strFlag, *boolFlag)
//}

func main() {

	argsLen := len(os.Args)
	args := os.Args
	logger.Info("args : ", args)
	logger.Info("args len : ", argsLen)

	if argsLen > 1 {
		arg1 := os.Args[1]
		fmt.Println("arg1 : ", arg1)
	}

	if argsLen > 2 {
		arg2 := os.Args[2]
		fmt.Println("arg2 : ", arg2)
	}

	fmt.Println("the end ...")
}
