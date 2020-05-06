package invert
import (
	"fmt"
	"os"
	"text/scanner"
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

func (index *InvertedIndex) Parse(filename string) {
	if index.alreadyParsed(filename) {
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		cwd, _ := os.Getwd()
		fmt.Printf("The file named %s doesn't exist, the current directory is %s\n",
			filename, cwd)
		return
	}
	defer file.Close()
	var s scanner.Scanner
	s.Init(file)
	var token rune
	for token != scanner.EOF {
		token = s.Scan()
		tokenText := s.TokenText()
		_, found := index.indexMap[tokenText]
		pos := s.Pos()
		result := Result{File: filename,
			Line:  pos.Line,
			Index: pos.Column - len(tokenText)}
		if !found {
			index.indexMap[tokenText] = []Result{result}
		} else {
			index.indexMap[tokenText] = append(index.indexMap[tokenText],
				result)
		}
	}
	index.files = append(index.files, filename)
}

// Get search the text in this inverted index
func (index *InvertedIndex) Get(text string) []Result {
	return index.indexMap[text]
}

func (index *InvertedIndex) alreadyParsed(filename string) bool {
	for _, file := range index.files {
		if file == filename {
			return true
		}
	}
	return false
}