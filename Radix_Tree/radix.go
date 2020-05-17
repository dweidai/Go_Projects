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

