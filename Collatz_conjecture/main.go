package main

import "fmt"

func main(){
	count :=0
	var n int
	fmt.Println("input a positive integer")
	_, err := fmt.Scanf("%d", &n)
	if err != nil{
		fmt.Println(err)
		return
	}
	if n <=0 {
		fmt.Println("input a positive integer")
		return
	}
	for n > 1{
		if n%2 == 0{
			n /= 2
		} else{
			n = n*3 +1
		}
		count ++
	}

	fmt.Println("Number of steps is ", count)
}