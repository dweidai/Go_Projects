package main

import tl "github.com/JoelOtter/termloop"


var score = 0
//var game *tl.game
//var scoreText *tl.Text
//var field *Field

func main(){
	game := tl.NewGame()
	/*main := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	width := 80
	height := 30
	field := NewField(width, height)
	*/
	game.Start()
}