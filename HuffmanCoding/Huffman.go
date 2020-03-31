type huffmanTree struct {
	nodes    []huffmanNode
	nextNode int
}

type huffmanNode struct {
	left, right           uint16
	leftValue, rightValue uint16
}
