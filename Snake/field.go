package main

import tl "github.com/JoelOtter/termloop"

type Node struct{
	x int
	y int
}

type Field struct{
	*tl.Entity
	width int
	height int
	edges map[Node]int
}

func newField(width int, height int) *Field{
	f := new(Field)
	f.Entity = tl.NewEntity(1, 1, 1, 1)
	f.width = width-1
	f.height = height-1
	f.edges = make(map[Node]int)
	for i := 0; i < f.width; i++ {
		f.edges[Node{x, 0}] = 1
		f.edges[Node{x, f.height}] = 1
	}
	for j:=0; j<f.height; j++{
		f.edges[Node{0, y}] = 1
		f.edges[Node{f.width, y}] = 1
	}
	return f
}

func (f *Field) update(screen *tl.Screen){
	for c:range f.edges{
		screen.RenderCell(c.x, c.y, &tl.Cell{
				Bg: tl.ColorBlue,
		})
	}
}

func (f *Field) contain(cord Node) bool{
	value, ok := f.edges[cord]
	if value != 0 && value != 1{
		return false
	} 
	return ok
}