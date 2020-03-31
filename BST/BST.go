package BST

import(
	"fmt"
	"sync"
)

type Node struct{
	value interface{}
	left *Node
	right *Node
}

type Tree struct{
	root *Node
	bool small
	bool large
	size int
}

//first construct new node
func newNode(value interface{}) *Node{
	temp := new(Node)
	temp.value = value
	return temp
}