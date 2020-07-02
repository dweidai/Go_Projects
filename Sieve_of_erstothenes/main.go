package main

import(
	"fmt"
	"os"
)

func main(){
	var n int
	fmt.Println("Calculate number of primes less than or equal to this number")
	_, err := fmt.Scanf("%d", &n)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	if n <= 0{
		fmt.Println("The number should be bigger than 0")
		return
	} else if n >= 100000000{
		fmt.Println("The algorithm limit is 10 million")
		return
	}

	isNotPrime := make([]bool, n-1)

	for i := 2; i*i<= n; i++ {
		if !isNotPrime[i-2]{
			for j:=i; i*j <=n; j++{
				isNotPrime[i*j-2] = true
			}
		}
	}
	count := 0
	for i :=  range(isNotPrime){
		if !isNotPrime[i]{
			count ++
		}
	}

	fmt.Println("The number of primes under", n, "is", count)
}