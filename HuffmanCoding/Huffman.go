package huffman

import(
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Node struct{
	Parent *Node
	Left *Node
	Right *Node
	Count int
	Value int
}


func MakeNodes(text string) map[string]int{
	m := make(map[string]int)
	dat, err := ioutil.ReadFile(text)
	data := string(dat)
	for i,c := range data{
		
	}
}