package main

import (
	"fmt"
	"go_study/tree/node"
	"strconv"
)

func buildTree() node.Node {
	var tree node.Node
	tree.Value = 0

	tree.LeftNode = new(node.Node)
	tree.LeftNode.Value = 1

	tree.LeftNode.LeftNode = new(node.Node)
	tree.LeftNode.LeftNode.Value = 3

	tree.LeftNode.LeftNode.LeftNode = new(node.Node)
	tree.LeftNode.LeftNode.LeftNode.Value = 5

	tree.LeftNode.LeftNode.RightNode = new(node.Node)
	tree.LeftNode.LeftNode.RightNode.Value = 6

	tree.RightNode = new(node.Node)
	tree.RightNode.Value = 2

	tree.RightNode.LeftNode = new(node.Node)
	tree.RightNode.LeftNode.Value = 4

	tree.RightNode.LeftNode.RightNode = new(node.Node)
	tree.RightNode.LeftNode.RightNode.Value = 7

	return tree
}

func main() {
	tree := buildTree()

	count := 0
	tree.Do(func(n *node.Node) {
		count++
		fmt.Println("count : " + strconv.Itoa(count))
	})

	outChan := tree.DoWithChan()

	maxValue := 0
	for c := range outChan {
		if maxValue < c.Value {
			maxValue = c.Value
		}
	}
	fmt.Println("max value is ", maxValue)

}
