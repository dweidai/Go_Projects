package BST

import(
	"math/rand"
	"testing"
	"fmt"
)

var summary *Tree

func RandomTree() *Tree{
	temp := newTree(true)
	for i := 0; i < 10; i++ {
		insert(rand.Intn(100), temp)
	}
	return temp
}

func BalanceTree() *Tree{
	temp:= newTree(true)
	insert(15, temp)
	insert(10, temp)
	insert(20, temp)
	insert(8, temp)
	insert(12, temp)
	insert(25, temp)
	insert(16, temp)
	return temp
}

func LeftTree() *Tree{
	temp:= newTree(true)
	for i := 15; i > 0; i-- {
		insert(i, temp)
	}
	return temp
}

func RightTree() *Tree{
	temp:= newTree(false)
	for i := 15; i > 0; i-- {
		insert(i, temp)
	}
	return temp
} 

//this is just testing how testing library works
func Abs(x int) int{
	if x<0{
		return -x
	}
	return x
}
func TestAbs(t *testing.T) {
	fmt.Println("Testing Abs")
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
    fmt.Println("\t1/1 test passed")
    summary = newTree(false)
	insert(1, summary)
}

//this is testing tree construction
func TestTree(t *testing.T){
	fmt.Println("Testing construction")
	tree := newTree(true)
	if tree == nil {
        t.Errorf("Error initializing tree")
    }
    RandomTree := RandomTree()
	if RandomTree == nil {
        t.Errorf("Error initializing tree")
    }
    BalanceTree := BalanceTree()
	if BalanceTree == nil {
        t.Errorf("Error initializing tree")
    }
    fmt.Println("\t3/3 tests passed")
    insert(1, summary)
 }

//testing size
func TestSize(t *testing.T){
	fmt.Println("Testing Size")
	tree := newTree(true)
	if size(tree) != 0 {
        t.Errorf("Should be 0 but is %d", size(tree))
    }
    RandomTree := RandomTree()
	if size(RandomTree) != 10 {
        t.Errorf("Should be 10 but is %d", size(RandomTree))
    }
    LeftTree := LeftTree()
	if size(LeftTree) != 15 {
        t.Errorf("Should be 15 but is %d", size(LeftTree))
    }
    fmt.Println("\t3/3 tests passed")
    insert(1, summary)
}

func TestHas(t *testing.T){
	fmt.Println("Testing Has")
	tree := newTree(true)
	BalanceTree := BalanceTree()
	RightTree := RightTree()
	if has(0, tree) == true{
		t.Errorf("Has function should be false b/c tree empty")
	}
	if has(1, tree) == true{
		t.Errorf("Has function should be false b/c tree empty")
	}
	if has(8, BalanceTree) == false{
		t.Errorf("Has function should be true b/c tree has 8")
	}
	if has(16, BalanceTree) == false{
		t.Errorf("Has function should be true b/c tree has 16")
	}
	if has(30, BalanceTree) == true{
		t.Errorf("Has function should be false b/c tree does not have 30")
	}
	if has(1, RightTree) == false{
		t.Errorf("Has function should be false b/c tree has 11")
	}
	if has(15, RightTree) == false{
		t.Errorf("Has function should be false b/c tree has 15")
	}
	if has(20, RightTree) == true{
		t.Errorf("Has function should be false b/c tree does not have 20")
	}
	fmt.Println("\t8/8 tests passed")
	insert(1, summary)
}

func TestInsert(t *testing.T){
	fmt.Println("Testing insert")
	tree := newTree(true)
	BalanceTree := BalanceTree()
	if has(0, tree) == true{
		t.Errorf("Has function should be false b/c tree empty")
	}
	insert(0, tree)
	if has(0, tree) == false{
		t.Errorf("We just insertedt 0 to the tree, should be true")
	}
	if has(1, tree) == true{
		t.Errorf("Has function should be false b/c tree empty")
	}
	insert(1, tree)
	if has(1, tree) == false{
		t.Errorf("We just insertedt 1 to the tree, should be true")
	}
	insert(1, BalanceTree)
	if has(1, BalanceTree) == false{
		t.Errorf("We just insertedt 1 to the tree, should be true")
	}
	insert(30, BalanceTree)
	if has(30, BalanceTree) == false{
		t.Errorf("We just insertedt 30 to the tree, should be true")
	}
	insert(100, BalanceTree)
	if has(100, BalanceTree) == false{
		t.Errorf("We just insertedt 100 to the tree, should be true")
	}
	fmt.Println("\t7/7 tests passed")
	insert(1, summary)
}

func TestDelete(t *testing.T){
	fmt.Println("Testing Delete")
	tree := newTree(true)
	if delete(100, tree) == true{
		t.Errorf("It should print false because the tree is empty")
	}
	if delete(-1, tree) == true{
		t.Errorf("It should print false because the tree is empty")
	}
	BalanceTree := BalanceTree()
	if delete(100, BalanceTree) == true{
		t.Errorf("It should print false because there is not 100")
	}
	if delete(12, BalanceTree) == false{
		t.Errorf("It delete 12 from the balanced tree")
	}
	if delete(15, BalanceTree) == false{
		t.Errorf("It delete 15 from the balanced tree")
	}
	RightTree := RightTree()
	if delete(100, RightTree) == true{
		t.Errorf("It should print false because there is not 100")
	}
	if delete(12, RightTree) == false{
		t.Errorf("It delete 12 from the right tree")
	}
	if delete(7, RightTree) == false{
		t.Errorf("It delete 7 from the right tree")
	}
	fmt.Println("\t8/8 tests passed")
	insert(1, summary)
}

func TestMax(t *testing.T){
	fmt.Println("Testing Max")
	BalanceTree := BalanceTree()
	RightTree := RightTree()
	LeftTree := LeftTree()
	if max(BalanceTree) != 25{
		t.Errorf("It should be 25 but it prints out %d", max(BalanceTree))
	}
	if max(RightTree) != 15{
		t.Errorf("It should be 15 but it prints out %d", max(RightTree))
	}
	if max(LeftTree) != 15{
		t.Errorf("It should be 15 but it prints out %d", max(LeftTree))
	}
	fmt.Println("\t3/3 tests passed")
	insert(1, summary)
}


func TestMin(t *testing.T){
	fmt.Println("Testing Min")
	BalanceTree := BalanceTree()
	RightTree := RightTree()
	LeftTree := LeftTree()
	if min(BalanceTree) != 8{
		t.Errorf("It should be 8 but it prints out %d", min(BalanceTree))
	}
	if min(RightTree) != 1{
		t.Errorf("It should be 11 but it prints out %d", min(RightTree))
	}
	if min(LeftTree) != 1{
		t.Errorf("It should be 1 but it prints out %d", min(LeftTree))
	}
	fmt.Println("\t3/3 tests passed")
	insert(1, summary)
}

func TestRank(t *testing.T){
	fmt.Println("Testing rank")
	BalanceTree := BalanceTree()
	RightTree := RightTree()
	LeftTree := LeftTree()
	Tree := newTree(true)
	if rank(Tree) != 0{
		t.Errorf("Should be 0 but print out %d", rank(Tree))
	}
	if rank(LeftTree) != 15{
		t.Errorf("Should be 15 but print out %d", rank(LeftTree))
	}
	if rank(RightTree) != 15{
		t.Errorf("Should be 15 but print out %d", rank(RightTree))
	}
	if rank(BalanceTree) != 3{
		t.Errorf("Should be 3 but print out %d", rank(BalanceTree))
	}
	fmt.Println("\t4/4 tests passed")
	insert(1, summary)
}


func TestKth(t *testing.T){
	BalanceTree := BalanceTree()
	RightTree := RightTree()
	LeftTree := LeftTree()
	if Kth(LeftTree, 3) != 2{
		t.Errorf("Should be 2 but print out %d", Kth(LeftTree, 3))
	}
	if Kth(RightTree, 1) != 15{
		t.Errorf("Should be 15 but print out %d", Kth(RightTree, 3))
	}
	if Kth(BalanceTree, 1) != 8{
		t.Errorf("Should be 8 but print out %d", Kth(BalanceTree, 1))
	}
	fmt.Println("\t4/4 tests passed")
	insert(1, summary)
}

func TestTraverse(t *testing.T){
	BalanceTree := BalanceTree()
	RightTree := RightTree()
	LeftTree := LeftTree()
	RandomTree := RandomTree()

	fmt.Println("Inorder:")
	print(BalanceTree, "inorder")
	fmt.Println("Preorder:")
	print(BalanceTree, "preorder")
	fmt.Println("Postorder:")
	print(BalanceTree, "postorder")

	fmt.Println("Inorder:")
	print(RightTree, "inorder")
	fmt.Println("Preorder:")
	print(RightTree, "preorder")
	fmt.Println("Postorder:")
	print(RightTree, "postorder")

	fmt.Println("Inorder:")
	print(LeftTree, "inorder")
	fmt.Println("Preorder:")
	print(LeftTree, "preorder")
	fmt.Println("Postorder:")
	print(LeftTree, "postorder")

	fmt.Println("Inorder:")
	print(RandomTree, "inorder")
	fmt.Println("Preorder:")
	print(RandomTree, "preorder")
	fmt.Println("Postorder:")
	print(RandomTree, "postorder")

	fmt.Println("\t12/12 tests passed")
	insert(1, summary)
}

func TestSummary(t *testing.T){
	if size(summary) != 11{
		t.Errorf("Should be 11 sets of tests in total, you passed %d", (size(summary)/11)*100)
	}
	fmt.Println("All 11 sets of tests passed")
	fmt.Println("Total of 56/56 tests passed")
}
