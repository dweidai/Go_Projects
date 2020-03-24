package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		count("sheep")
		wg.Done()
	}()
	//go count("sheep")
	//go count("sticks")
	
	//time.Sleep(time.Millisecond * 2000)
	//fmt.Scanln()

	wg.Wait()
}

func count(thing string){
	for i:=1; i<=10; i++{
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

}