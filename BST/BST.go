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
	bool small // to keep track how is the tree structured
	size int
}

//first construct new node
func newNode(value interface{}) *Node{
	temp := new(Node)
	temp.value = value
	return temp
}

func newTree(small bool) *Tree{
	temp := new(Tree)
	temp.small = small
	return temp
}

func size(tree *Tree) int{
	return tree.size
}

func insert(value interface{}, tree *Tree){
	if size(tree) == 0{
		tree.root = newNode(value)
		tree.size ++
	}else{
		helper_insert(value, tree.root, tree.small)
		tree.size ++
	}
}

func helper_insert(value interface{}, node *Node, bool small){
	if small == true{
		if value < node.value{
			if node.left == nil{
				node.left = newNode(value)
			} else {
				helper_insert(value, node.left, small)
			}
		} else{
			if node.right == nil{
				node.right = newNode(value)
			} else {
				helper_insert(value, node.right, small)
			}
		}
	}else{
		if value > node.value{
			if node.left == nil{
				node.left = newNode(value)
			} else{
				helper_insert(value, node.left, small)
			}
		} else{
			if node.right == nil{
				node.right = newNode(value)
			} else {
				helper_insert(value, node.right, small)
			}
		}
	}
}

func preorder(tree *Tree){

}