package main

import(
	"fmt"
	"math/rand"
)

var grid[3][3] int
var user1 int
var user2 int

func main() {
	fmt.Println("\tGame Initializing")
	if rand.Intn(2) != -1{
		fmt.Println("User1 Start")
		user1 = 1
		user2 = -1

	} else{
		fmt.Println("User2 Start")
		user1 = -1
		user2 = 1
	}
	//fmt.Println(detect())
	var x int
	var y int
	count := 1
	for detect() != 1{
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
		if count > 0{
			count -= 2
		}else if count < 0{
			count +=2
		}
		placekey(x,y,count)
		fmt.Println(detect())
		print()
	}

	/*print()
	placekey(0,0,user1)
	placekey(1,1,user2)
	placekey(2,2,user1)
	placekey(0,0,user2)
	print()*/
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
