pacakge main

import(

)

type Field struct{
	cells [][]int
	width int
	height int
}

type Snake struct{
	x int
	y int
	length int
}

var snake *Snake
var field *Field

