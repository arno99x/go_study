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

func (node Node) DoWithChan() chan *Node {
	out := make(chan *Node)
	go func() {
		node.Do(func(n *Node) {
			out <- n
		})
		close(out)
	}()

	return out
}
