package main

import(
	"fmt"
	"math/rand"
)

var grid[3][3] int
var user1 int
var user2 int
var situation int
const MAX = 9

func main() {
	fmt.Println("\tGame Initializing")
	print()

	if rand.Intn(2) == 0{
		fmt.Println("User1 Start")
		user1 = 1
		user2 = -1
		situation = 1

	} else{
		fmt.Println("User2 Start")
		user1 = -1
		user2 = 1
		situation = 0
	}
	//fmt.Println(detect())
	var x int
	var y int
	count := 1
	step :=0
	for detect() != 1{
		step++
		fmt.Println("Please Input the wanted (row column)")
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
		if count > 0{
			count -= 2
		}else if count < 0{
			count +=2
		}
		for placekey(x,y,count)!= 0{
			fmt.Println("Please Re-Input the wanted (row column)")
			fmt.Scanf("%d", &x)
			fmt.Scanf("%d", &y)
		}
		print()
		if detect() == 1{
			if situation ==1 && count == 1{
				fmt.Println("User1 WIN!!!")
			}else if situation ==1 && count == -1{
				fmt.Println("User2 WIN!!!")
			}
			if situation ==0 && count == -1{
				fmt.Println("User1 WIN!!!")
			}else if situation ==0 && count == 1{
				fmt.Println("User2 WIN!!!")
			}
		}
		if(step == MAX){
			fmt.Println("It's a TIE!!!")
			break
		}
	}

}

func placekey(x int, y int, z int) int{
	if x<0 || x >2 || y<0 || y>2{
		fmt.Errorf("invalid index")
		return -1
	}else if grid[x][y] != 0{
		fmt.Errorf("Already placed")
		return 1
	}else{
		grid[x][y] = z
		return 0
	}
}

func print(){
	fmt.Println("\nCurrent Status:")
	for i:=0; i<3; i++{
		fmt.Println("-------")
		for j:=0; j<3; j++{
			if j==0{
				fmt.Print("|")
			}
			if grid[i][j] == 1{
				fmt.Print("X|")
			} else if grid[i][j] == -1{
				fmt.Print("O|")
			} else{
				fmt.Print(" |")
			}
			
		}
		fmt.Println()
	}
	fmt.Println("-------")
}

func detect() int{
	for i := 0; i < 3; i++ {
		if grid[i][1] == grid[i][0] && grid[i][1] == grid[i][2] && grid[i][1] != 0{
			return 1
		}
		if grid[1][i] == grid[0][i] && grid[1][i] == grid[2][i] && grid[1][i] != 0{
			return 1
		}
	}
	if grid[1][1] == grid[0][0] && grid[1][1]== grid[2][2] && grid[0][0] != 0{
		return 1
	}
	if grid[0][2] == grid[1][1] && grid[0][2] == grid[2][0] && grid[1][1] != 0{
		return 1
	}
	return 0
	
}
