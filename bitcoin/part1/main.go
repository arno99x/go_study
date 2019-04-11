package main

import (
	"fmt"
	"go_study/bitcoin/part1/core"
)

func main() {
	bc := core.NewBlockChain()
	bc.AddBlock("send 1 BTC to arno")
	bc.AddBlock("send 2 more BTC to arno")

	for _, block := range bc.Blocks {
		fmt.Printf("prev.hash : %x\n", block.PrevBlockHash)
		fmt.Printf("hash : %x\n", block.Hash)
		fmt.Printf("data : %s\n\n", block.Data)
	}
}
