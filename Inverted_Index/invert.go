package invert
import (
	"fmt"
	"os"
)

type InvertedIndex struct {
	indexMap map[string][]Result
	files	[]string
}

type Result struct{
	File string
	line int
	index int
}

func newInvertedIndex() *InvertedIndex{
	index := &InvertedIndex()
	index.indexMap = make(map[string][]Result)
	return index
}