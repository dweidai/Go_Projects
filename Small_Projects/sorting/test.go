package sort

import(
	"testing"
	"fmt"
)

func TestBubbleSort(t *testing.T){
	fmt.Println("Testing bubble sort")
	temp := []int{11, 13, 5, 37, 41, 3, 2, 7, 17}
	answer := []int{2, 3, 5, 7, 11, 13, 17, 37, 41}
	array := bubblesort(temp)
	for i:=0; i<len(temp); i++{
		if array[i] != answer[i]{
			t.Errorf("Error bubblesort")
		}
    }
    if len(answer) != len(array) || len(temp) != len(array) {
    	t.Errorf("Error bubblesort")
    }
    temp = []int{11, 13, 5}
	answer = []int{5, 11, 13}
	array = bubblesort(temp)
	for i:=0; i<len(temp); i++{
		if array[i] != answer[i]{
			t.Errorf("Error bubblesort")
		}
    }
     if len(answer) != len(array) || len(temp) != len(array) {
    	t.Errorf("Error bubblesort")
    }
    temp = []int{11, 11, 11, 11, 11}
	answer = []int{11, 11, 11, 11, 11}
	array = bubblesort(temp)
	for i:=0; i<len(temp); i++{
		if array[i] != answer[i]{
			t.Errorf("Error bubblesort")
		}
    }
     if len(answer) != len(array) || len(temp) != len(array) {
    	t.Errorf("Error bubblesort")
    }
    temp = []int{1}
	answer = []int{1}
	array = bubblesort(temp)
	for i:=0; i<len(temp); i++{
		if array[i] != answer[i]{
			t.Errorf("Error bubblesort")
		}
    }
     if len(answer) != len(array) || len(temp) != len(array) {
    	t.Errorf("Error bubblesort")
    }

    fmt.Println("\t4/4 tests passed")
    
 } 

 func TestMergeSort(t *testing.T){
	fmt.Println("Testing merge sort")
	temp := []int{11, 13, 5, 37, 41, 3, 2, 7, 17}
	answer := []int{2, 3, 5, 7, 11, 13, 17, 37, 41}
	array := mergesort(temp)
	for i:=0; i<len(temp); i++{
		if array[i] != answer[i]{
			t.Errorf("Error mergesort")
		}
    }
    if len(answer) != len(array) || len(temp) != len(array) {
    	t.Errorf("Error mergesort")
    }
    temp = []int{11, 13, 5}
	answer = []int{5, 11, 13}
	array = mergesort(temp)
	for i:=0; i<len(temp); i++{
		if array[i] != answer[i]{
			t.Errorf("Error mergesort")
		}
    }
     if len(answer) != len(array) || len(temp) != len(array) {
    	t.Errorf("Error mergesort")
    }
    temp = []int{11, 11, 11, 11, 11}
	answer = []int{11, 11, 11, 11, 11}
	array = mergesort(temp)
	for i:=0; i<len(temp); i++{
		if array[i] != answer[i]{
			t.Errorf("Error mergesort")
		}
    }
     if len(answer) != len(array) || len(temp) != len(array) {
    	t.Errorf("Error mergesort")
    }
    temp = []int{1}
	answer = []int{1}
	array = mergesort(temp)
	for i:=0; i<len(temp); i++{
		if array[i] != answer[i]{
			t.Errorf("Error mergesort")
		}
    }
     if len(answer) != len(array) || len(temp) != len(array) {
    	t.Errorf("Error mergesort")
    }

    fmt.Println("\t4/4 tests passed")
    
 }