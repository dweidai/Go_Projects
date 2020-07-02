package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Enter a number to search for the next prime")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	var text string 
	if isPrime(num){
		num += 1
	}
	count := 0
	for count == 0 {
		for isPrime(num) == false {
			num += 1
		}
		fmt.Println("Current: ", num)
		fmt.Println("Enter \"next\" or \"n\" to find the next prime")
		_, err := fmt.Scan(&text)
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		if text == "next" || text == "n"{
			num += 1
		} else{
			count = 1
		}
	}
	fmt.Println("Final Prime: ", num)

	
}

func isPrime(n int) bool{
	if n <= 1{
		return false
	}
	for i:=2; i<n; i++{
		if n%i == 0{
			return false
		}
	}
	return true
}