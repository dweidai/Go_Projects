package main

import(
	"fmt"
	"os"
	"math/big"
)

func fact(a int) *big.Int {
	b := big.NewInt(1)
	for i := a; i > 1; i-- {
		b = b.Mul(b, big.NewInt(int64(i)))
	}
	return b
}

func piNth(num int){
	pi := big.NewRat(0, 1)
	temp := big.NewInt(1)
	term := big.NewRat(0.0, 1.0)
	for i:=1; i<4*num+10; i++{
		top := big.NewInt(int64(i))
		temp = temp.Exp(big.NewInt(2), big.NewInt(int64(i)), nil)
		top = top.Mul(top, temp)
		temp = temp.Exp(temp.Set(fact(i)), big.NewInt(2), nil)
		top = top.Mul(top, temp)
		bottom := fact(2 * i)
		term := term.SetFrac(top, bottom)
		pi.Add(pi, term)
	}
	pi = pi.Add(pi, big.NewRat(-3, 1))
	fmt.Println("π to", num, "decimal places: ")
	fmt.Println(pi.FloatString(num))
}

func main(){
	fmt.Println("enter a number less than 1000")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	if num > 1000 || num < 0{
		fmt.Println("Invalid user input range, should between 0 and 1000")
		os.Exit(1)
	}
	piNth(num)
}