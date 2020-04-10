package main

import (
	"math/rand"
	tl "github.com/JoelOtter/termloop"
)

type food struct{
	*tl.Entity
	location Node
}

func newFood() *Food{
	f := new(Food)
	f.Entity = t1.NewEntity(1,1,1,1)
	f.location = new(Node)
	f.location.x := randInRange(1, border.width-1)
	f.location.y := randInRange(1, border.height-1)
	return f
}

func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func (f *Food)position()(int, int){
	return f.node.x, f.node.f
}

func (f *Food)display(screen *tl.Screen){
	screen.RenderCell(f.coord.x, f.coord.y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '*',
	})
}

func (f *Food) snakeCollision() {
	f = newFood()
}