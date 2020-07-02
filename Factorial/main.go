package main

import "fmt"

func fac_loop(n int) int{
	if n==0{
		return 1
	}
	result := n
	for i := 1; i < n; i++ {
		result = result * i
		
	}
	return result
}

func fac_rec(n int) int{
	if n == 0{
		return 1
	}
	return n * fac_rec(n-1)
}

func main(){
	var n int
	fmt.Println("input an integer")
	fmt.Scanf("%d", &n)
	fmt.Println("Loop implementation: ", fac_loop(n))
	fmt.Println("Recursive implementation: ", fac_rec(n))
}