package main

import tl "github.com/JoelOtter/termloop"

var game *tl.game
var field *Field

func endGame(){
	end := tl.NewBaseLevel(tl.Cell{
		Bg:tl.ColorBlack,
	})
	game.Screen().SetLevel(end)
}

func main(){
	game := tl.NewGame()
	main := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	width := 80
	height := 30
	field := NewField(width, height)
	snake := NewSnake()
	food := NewFood()

	main.AddEntity(field)
	main.AddEntity(snake)
	main.AddEntity(food)

	game.Screen.SetLevel(main)
	game.Screen.SetFps(10)
	game.Start()
}