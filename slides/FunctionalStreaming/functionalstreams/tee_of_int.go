package functionalstreams

func (node *NodeOfInt) AddTee(firstNode, secondNode *NodeOfInt) *NodeOfInt {
	node_A, node_B := firstNode.Tee()
	node_A.Connect(secondNode)
	node_B.Connect(node)
	return node
}
