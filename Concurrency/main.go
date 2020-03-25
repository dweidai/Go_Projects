package main

import (
	"fmt"
//	"time"
//	"sync"
)

func fib(n int) int{
	if n<=1{
		return n
	}
	 return fib(n-1) + fib(n-2)
}

func worker(jobs <-chan int, results chan<- int){
	for n:= range jobs{
		results <- fib(n)
	}
}
func main(){

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)


	for i:=0; i< 100; i++ {
		jobs <- i
	}

	close(jobs)
	for j:=0; j<100; j++ {
		fmt.Println(<-results)
	}
	/*c := make(chan string, 2)
	c <- "Hello"
	c <- "Baby"

	msg := <- c
	fmt.Println(msg)
	msg = <- c
	fmt.Println(msg)
	c1 := make(chan string)
	c2 := make(chan string)
	go func(){
		for{
			c1 <-"Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func(){
		for{
			c2 <- "Every two seconds"
			time.Sleep(time.Millisecond*2000)
		}
	}()

	for{
		select{
		case msg := <- c1:
			fmt.Println(msg)
		case mss := <- c2:
			fmt.Println(mss)
		}
	}*/
}

/*
func main(){
	//var wg sync.WaitGroup //this is just a counter
	//wg.Add(1)

	//go func(){
	//	count("sheep")
	//	wg.Done()
	//}()
	c := make(chan string)
	go count("sheep", c)
	//go count("sticks")
	for msg := range c{
		fmt.Println(msg)
	}
	
	//time.Sleep(time.Millisecond * 2000)
	//fmt.Scanln() //easiest fix but not propoer

	//wg.Wait()
	//close(c) //should cause panic because close the channel prematurely
}

func count(thing string, c chan string){
	for i:=1; i<=10; i++{
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c) //right place for closing the channel

}
*/