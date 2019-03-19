package node

type Node struct {
	LeftNode  *Node
	RightNode *Node
	Value     int
}

func (node Node) Do(f func(n *Node)) {
	f(&node)
	if node.LeftNode != nil {
		node.LeftNode.Do(f)
	}
	if node.RightNode != nil {
		node.RightNode.Do(f)
	}
}
