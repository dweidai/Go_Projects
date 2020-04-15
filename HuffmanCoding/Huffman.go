package main

import(
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	)

type Node struct{
	Parent *Node
	Left *Node
	Right *Node
	Count int
	Value string
}

/*
 * This method is used to load the files in to the map
 * And we can count the number of characters appeared
 * And then easier to load in to the huffman nodes
 */
func MakeNodes(text string) []*Node{
	m := make(map[string]int)
	dat, err := ioutil.ReadFile(text)
	if err != nil{
		fmt.Println(err)
	}
	data := string(dat)
	for i, c := range data{
		if i <0{
			fmt.Println("INDEX INCORRECT")
		}
		m[string(c)] = m[string(c)]+1
	}
	toReturn := make([]*Node, int(len(m)))
	count := 0
	for key, element:= range m{
		temp := new(Node)
		temp.Count = element
		temp.Value = key
		toReturn[count] = temp
		count++
	}
	return toReturn
}

type SortNodes []*Node
func (sn SortNodes) Len() int           { return len(sn) }
func (sn SortNodes) Less(i, j int) bool { return sn[i].Count < sn[j].Count }
func (sn SortNodes) Swap(i, j int)      { sn[i], sn[j] = sn[j], sn[i] }


//code returns the huffman code of the node
// left children get 0 and right get 1 and then
// we will traverse up to keep track of the numbers
func (n *Node) Encode() (r uint64, bits byte){
	for parent := n.Parent; parent != nil; n, parent = parent, parent.Parent{
		if parent.Right == n{
			r |= 1<< bits
		}
		bits ++
	}
	return
}

func BuildTree(leaves []*Node) *Node{
	if len(leaves) == 0{
		return nil
	} else{
		sort.Stable(SortNodes(leaves))
	}
	length := len(leaves)
	for length > 1{
	 	left := leaves[0]
	 	right := leaves[1]
	 	pCount := left.Count + right.Count
	 	parent := new(Node)
	 	parent.Left = left
	 	parent.Right = right
	 	parent.Count = pCount
	 	left.Parent = parent
	 	right.Parent = parent

	 	ls := leaves[2:]
	 	index := sort.Search(len(ls), func (i int) bool { return ls[i].Count >= pCount})
	 	index += 2

	 	copy(leaves[1:], leaves[2:index])
	 	leaves[index-1] = parent

	 	leaves = leaves[1:]
	 	length = len(leaves)
	 }
	 return leaves[0]
}


func Print(root *Node){
	var traverse func(n *Node, code uint64, bits byte)

	traverse = func(n *Node, code uint64, bits byte){
		if n.Left ==nil{
			fmt.Printf("'%c': %0"+strconv.Itoa(int(bits))+"b\n", n.Value, code)
			return
		}
		bits++
		traverse(n.Left, code<<1, bits)
		traverse(n.Right, code<<1+1, bits)
	}

	traverse(root, 0, 0)
}



func main(){
	Print(BuildTree(MakeNodes("./TEST.TXT")))
	Print(BuildTree(MakeNodes("./TEXT.txt")))
}
