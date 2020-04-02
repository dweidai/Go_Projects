package BST

import(
	"fmt"
	"math"
	"strings"
)

// these are structs of the node and the tree
type Node struct{
	value int
	left *Node
	right *Node
}

// bool small is to keep track how is the tree structured
type Tree struct{
	root *Node
	size int
	small bool
}

//first construct new node
func newNode(value int) *Node{
	temp := new(Node)
	temp.value = value
	return temp
}

//constructor for a new tree
func newTree(small bool) *Tree{
	temp := new(Tree)
	temp.small = small
	temp.size = 0
	return temp
}

//get the size of the tree
func size(tree *Tree) int{
	return tree.size
}

//insert a value into the tree
func insert(value int, tree *Tree) bool{
	if size(tree) == 0{
		tree.root = newNode(value)
		tree.size ++
		return true
	}else{
		if helper_insert(value, tree.root, tree.small){
			tree.size ++
			return true
		}
		return false
	}
}

//helper method with recursively insert the value to the right location
func helper_insert(value int, node *Node, small bool) bool{
	if small == true{
		if value < node.value{
			if node.left == nil{
				node.left = newNode(value)
				return true
			} else {
				return helper_insert(value, node.left, small)
			}
		} else{
			if node.right == nil{
				node.right = newNode(value)
				return true
			} else {
				return helper_insert(value, node.right, small)
			}
		}
	}else{
		if value > node.value{
			if node.left == nil{
				node.left = newNode(value)
				return true
			} else{
				return helper_insert(value, node.left, small)
			}
		} else{
			if node.right == nil{
				node.right = newNode(value)
				return true
			} else {
				return helper_insert(value, node.right, small)
			}
		}
	}
}

//delete a value from the tree gvien the value
func delete(value int, tree *Tree) bool{
	if size(tree) ==  0{
		return false
	} else{
		if helper_delete(value, tree.root, tree.small){
			tree.size --
			return true
		}
		return false
	}
}

//recursive hlper method to delete the given value from the node traverse
func helper_delete(value int, node *Node, small bool) *Node{
	if node == nil{
		return nil;
	}
	if small == true{
		if value < node.value{
			node.left = helper_delete(value, node.left, small)
		} else if value > node.value{
			node.right = helper_delete(value, node.right, small)
		} else{
			if node.left == nil{
				return node.right
			} else if node.right == nil{
				return node.left
			}

			temp := findNext(node.right)
			node.value = temp.value
			node.right = helper_delete(node.right, node.value, small)
		}
		return node
	} else {
		if value > node.value{
			node.left = helper_delete(value, node.left, small)
		} else if value < node.value{
			node.right = helper_delete(value, node.right, small)
		} else{
			if node.left == nil{
				return node.right
			} else if node.right == nil{
				return node.left
			}

			temp := findNext(node.right)
			node.value = temp.value
			node.right = helper_delete(node.right, node.value, small)
		}
		return node
	}
	
}

//depending on small is true or not
//this is finding the next value to rotate up
func findNext(node *Node) *Node{
	for node.left != nil{
		node = node.left;
	}
	return node;

}


//need function to check if a given value is in the tree or not
func has(value int, tree *Tree) bool{
	if size(tree) == 0{
		return false
	} else if value == tree.root.value{
		return true
	} else{
		return helper_has(value, tree.root, tree.small)
	}
}

//quick helper function to find if a value is in from the traversed node
func helper_has(value int, node *Node, small bool) bool{
	if node == nil{
		return false
	}
	if small == true{
		if value < node.value{
			return helper_has(value, node.left, small)
		} else if value > node.value{
			return helper_has(value, node.right, small)
		} else{
			return true
		}
	} else{
		if value > node.value{
			return helper_has(value, node.left, small)
		} else if value < node.value{
			return helper_has(value, node.right, small)
		} else{
			return true
		}
	}
}

//find the minimum from the tree
func min(tree *Tree) int{
	return helper_min(tree.root, tree.small)
}

func helper_min(node *Node, small bool) int{
	if small{
		for node.left != nil{
			node = node.left;
		}
		return node.value;
	}else{
		for node.right != nil{
			node = node.right;
		}
		return node.value;
	}
}

// and we also need to find the max from the tree
func max(tree *Tree) int{
	return helper_max(tree.root, tree.small)
}

func helper_max(node *Node, small bool) int{
	if small{
		for node.right != nil{
			node = node.right;
		}
		return node.value;
	}else{
		for node.left != nil{
			node = node.left;
		}
		return node.value;
	}
}


//depth of the tree or the rank of the tree
//here we are going to say that the root has rank 1
func rank(tree *Tree) int{
	return helper_rank(tree.root)
}

//helper method to find the rank of the tree as described above
func helper_rank(node *Node) int{
	if node == nil{
		return 0
	}
	return 1+math.max(helper_rank(node.left), helper_rank(node.right))
}

//now we are going to print out the tree
func print(tree *Tree, order string){
	helper_print(tree.root, order)
}

func helper_print(node *Node, order string){
	if strings.ToLower(order) == "inorder"{
		inorder(node)
	}
	if strings.ToLower(order) == "preorder"{
		preorder(node)
	}
	if strings.ToLower(order) == "postorder"{
		postorder(node)
	}
}

func inorder(node *Node){
	if node == nil{
		return;
	}
	inorder(node.left)
	fmt.Print(node.value)
	fmt.Print(" ")
	inorder(node.right)
}

func preorder(node *Node){
	if node == nil{
		return;
	}
	fmt.Print(node.value)
	fmt.Print(" ")
	preorder(node.left)
	preorder(node.right)
}

func postorder(node *Node){
	if node == nil{
		return;
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Print(node.value)
	fmt.Print(" ")
}

//leetcode 230 find the Kth smallest element in a BST
// or find the Kth largest element in the BST depending on the small bool setting
func Kth(tree *Tree, k int) int{
	return helper_Kth(tree.root, k)
}

func helper_Kth(node *Node, k int) int{
	size := countNodes(node)
	if k <= size {
		return helper_Kth(node.left, k)
	}else if k > size+1 {
		return helper_Kth(node.right, k-1-size)
	}
	return node.value
}

//count the number of nodes below this node
func countNodes(node *Node) int{
	if node == nil{
		return 0
	}
	return 1 + countNodes(node.left) + countNodes(node.right)
}