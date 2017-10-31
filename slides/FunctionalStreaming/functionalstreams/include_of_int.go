package functionalstreams

func (nodeFrom *NodeOfInt) Include(nodeTo *NodeOfInt, nameOfNewNode string) *NodeOfInt {
	include_node := NewNodeOfInt(nameOfNewNode)
	nodeFrom.Connect(include_node)

	include_node.cf <- func(i int) int { return i * 10 * 100 }
	return include_node
}
