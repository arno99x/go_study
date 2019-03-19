package main

import (
	"go_study/005_function/treenode/node"
)

func BuildTree() node.Node {
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
	tree.Do()
}
