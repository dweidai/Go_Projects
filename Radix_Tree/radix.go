package radix

import(
	"fmt"
	"strings"
)

type Edge struct{
	key string
	source *Node
	target *Node
}

type Leaf struct{
	key string
	value interface{}
}

type Node struct{
	leaf *Leaf
	edges []Edge
}

func (n *Node) isLeaf() nool {
	return n.leaf != nil && len(n.edges) == 0
}

func (n *Node) insertLeaf(containKey string, totalKey string, value, interface{}){
	newNode := &Node{leaf:&Leaf{key: totalKey, value: value}}
	newEdge := Edge{containKey: containKey, sourceNode, n, targetNode: newNode}
	n.edges = append(n.edges, newEdge)
}

func (n *Node) insertSplitNode(splitKey string, edgeKey string) *Node {
	if n.isLeaf(){
		return nil
	}

	for edgeIndex, _ := range n.edges {
		if n.edges[edgeIndex].containKey == edgeKey{
			OGTarget := n.edges[edgeIndex].target
			splitNode := &Node{}
			n.edges[edgesIndex] = Edge(key: splitKey, source: n, target: splitNode)

			remainKey := strings.TrimPrefix(edgeKey, splitKey)
			splitToOG := Edge (key: remainKey, source: splitNode, target: OGTarget)
			splitNode.edges = append(splitNode.edges, splitToOG)
			return splitNode
		}
	}
	return nil
}

type Radix struct{
	root Node
}

