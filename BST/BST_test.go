package BST

import(
	"fmt"
	"math/rand"
	"testing"
)

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

//this is just testing how testing library works
func Abs(x int) int{
	fmt.Println("testing")
	if x<0{
		return -x
	}
	return x
}
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}

func TestTree(t *testing.T){
	tree = newTree(true)
	if tree == nil {
        t.Errorf("Error initializing tree")
    }
}