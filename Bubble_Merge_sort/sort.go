package sort


func bubblesort(input []int) ([]int){
	toReturn := make ([]int, len(input))
	for i := 0; i < len(input); i++ {
		toReturn[i] = input[i]		
	}
	var temp int
	for i:=0; i<len(toReturn)-1; i++{
		for j:=0; j<len(toReturn)-1; j++{
			if toReturn[j] > toReturn[j+1]{
				temp = toReturn[j+1]
				toReturn[j+1] = toReturn[j]
				toReturn[j] = temp
			}
		}
	}
	return toReturn
}

func mergesort(input []int) ([]int){
	temp := make([]int, len(input), cap(input))
	split(input, temp)
	return input
}

func split(array1 []int, array2 []int){
	if len(array1) <= 1 {
		return
	}else{
		mid := len(array1) / 2
		split(array1[:mid], array2[:mid])
		split(array1[mid:], array2[mid:])
		merge(array1, array2, mid)
		copy(array1, array2)
	}
}

func merge(array1 []int, array2 []int, mid int){
	iBegin := 0
    iMid := len(array1) / 2
    iEnd := len(array1)
    for j := range array1 {
        if (iBegin < mid && (iMid >= iEnd || array1[iBegin] <= array1[iMid])) {
            array2[j] = array1[iBegin]  
            iBegin++  
        } else {
            array2[j] = array1[iMid]
            iMid++  
        }
    }
}