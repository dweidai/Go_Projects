package main

import(
	"math"
	"fmt"
	"os"
)

func main(){
	var price float64
	var payment float64

	fmt.Println("What is the price?")
	_, err := fmt.Scanf("%f", &price)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("What is the payment amount?")
	_, err = fmt.Scanf("%f", &payment)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	if price > payment{
		fmt.Println("Not enough gold for the purchase")
		return
	}
	price = math.Floor(price*100)/100
	payment = math.Floor(payment*100)/100
	diff := math.Floor((payment-price)*100)/100

	hundreds := math.Floor(diff / 100)
	leftover := diff - hundreds*100
	fifties := math.Floor(leftover / 50)
	leftover -= fifties * 50
	twenties := math.Floor(leftover / 20)
	leftover -= twenties * 20
	tens := math.Floor(leftover / 10)
	leftover -= tens * 10
	fives := math.Floor(leftover / 5)
	leftover -= fives * 5
	ones := math.Floor(leftover)
	leftover -= ones

	leftover *= 100
	quarters := math.Floor(leftover / 25)
	leftover -= quarters * 25
	dimes := math.Floor(leftover / 10)
	leftover -= dimes * 10
	nickels := math.Floor(leftover / 5)
	leftover -= nickels * 5
	pennies := math.Floor(leftover + 0.5)

	if hundreds > 0 {
		fmt.Println("$100 bills:", hundreds)
	}
	if fifties > 0 {
		fmt.Println("$50 bills: ", fifties)
	}
	if twenties > 0 {
		fmt.Println("$20 bills: ", twenties)
	}
	if tens > 0 {
		fmt.Println("$10 bills: ", tens)
	}
	if fives > 0 {
		fmt.Println("$5 bills:  ", fives)
	}
	if ones > 0 {
		fmt.Println("$1 bills:  ", ones)
	}
	if quarters > 0 {
		fmt.Println("Quarters:  ", quarters)
	}
	if dimes > 0 {
		fmt.Println("Dimes:     ", dimes)
	}
	if nickels > 0 {
		fmt.Println("Nickels:   ", nickels)
	}
	if pennies > 0 {
		fmt.Println("Pennies:   ", pennies)
	}
}