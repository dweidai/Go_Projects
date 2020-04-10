package main

import tl "github.com/JoelOtter/termloop"

type Cord struct{
	x int
	y int
}

type Field struct{
	*tl.Entity
	width int
	height int
	cords map[Cord]int
}

func NewField(width int, height int) *Field{
	f := Field
	f.Entity = tl.NewEntity(1, 1, 1, 1)
	f.width = width-1
	f.height = height-1
	f.cords = make(map[Cord]int)
	for i := 0; i < f.width; i++ {
		f.cords[Cord{x, 0}] = 1
		f.cords[Cord{x, f.height}] = 1
	}
	for j:=0; j<f.height; j++{
		f.cords[Cord{0, y}] = 1
		f.cords[Cord{f.width, y}] = 1
	}
	return f
}

func (f *Field) Draw(screen *tl.Screen){
	for c:range f.cords{
		screen.RenderCell(c.x, c.y, &tl.Cell{
				Bg: tl.ColorBlue,
		})
	}
}