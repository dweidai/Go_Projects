package main

import "fmt"

func main(){
	fmt.Println("Enter a number to generate the Fibonacci sequence to that number")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil{
		fmt.Println(err)
		return
	}
	if num < 1 || num > 92{
		fmt.Println("Computing limit is between 1 and 92")
	}
	fmt.Println("Fibonacci sequence to", num)
	if num == 1{
		fmt.Println("0")
	} else if num == 2{
		fmt.Println("0 1")
	} else{
		fmt.Print("0 1 ")
		p := 0
		q := 1
		count := 2
		for count <= num{
			if p<q{
				p += q
				fmt.Print(p)
			} else{
				q += p
				fmt.Print(q)
			}
			fmt.Print(" ")
			count += 1
		}
	}
	fmt.Println(" \n")
}