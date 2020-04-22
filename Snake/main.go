package main

import (
	"fmt"
	"math/rand"
	"time"
	 "os"
	tl "github.com/JoelOtter/termloop"
)

var width int
var height int

const(
	right direction  =  iota
	left
	up 
	down 
)

type direction int

type Node struct{
	x int
	y int
}

type Food struct{
	*tl.Entity
	location Node
}

func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func newFood() *Food{
	f := new(Food)
	f.Entity = tl.NewEntity(1,1,1,1)
	width = 80
	height = 30
	newX := randInRange(1, width-1)
	newY := randInRange(1, height-1)
	f.location.x, f.location.y = newX, newY
	f.SetPosition(newX, newY)
	return f
}

func (f *Food) Draw(screen *tl.Screen){
	screen.RenderCell(f.location.x, f.location.y, &tl.Cell{
		Fg: tl.ColorBlack,
		Ch: '@',
	})
}

func (f *Food) Collide(collision tl.Physical){
	switch collision.(type){
	case *Snake:
		newX := randInRange(1, width-1)
		newY := randInRange(1, height-1)
		f.location.x, f.location.y = newX, newY
		f.SetPosition(newX, newY)
	}
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
		f.edges[Node{i, 0}] = 1
		f.edges[Node{i, f.height}] = 1
	}
	for j:=0; j<f.height +1; j++{
		f.edges[Node{0, j}] = 1
		f.edges[Node{f.width, j}] = 1
	}
	return f
}

func (f *Field) Draw(screen *tl.Screen){
	for c := range f.edges{
		screen.RenderCell(c.x, c.y, &tl.Cell{
				Bg: tl.ColorBlack,
		})
	}
}

func (f *Field) Contains(cord Node) bool{
	_, ok := f.edges[cord]
	return ok
}


type Snake struct{
	*tl.Entity
	body []Node
	direction direction
	length int
	field *Field
}

func newSnake(field *Field) *Snake{
	s := new(Snake)
	s.Entity = tl.NewEntity(7,7,1,1)
	s.body = []Node{ 
		{7, 10},
		{8, 10},
		{9, 10},
	}
	s.length = len(s.body)
	s.direction = right
	s.field = field
	return s
}

func (s *Snake) getHead() *Node{
	return &s.body[len(s.body)-1]
}

func (s *Snake) collideTest() bool{
	//fmt.Println("reached")
	for i:=0; i<len(s.body)-1; i++{
		if *s.getHead() == s.body[i]{
			return true
		}
	}
	if s.field.Contains(*s.getHead()){
		fmt.Println("Head crush to field")
		return true
	}
	return false
}

func (s *Snake) Collide(collision tl.Physical){
	switch collision.(type){
	case *Food:
		//fmt.Println("FOOD!")
		s.length = s.length + 3
	case *Field:
		os.Exit(1)
	}
}


func (s *Snake) Draw(screen *tl.Screen){
	newHead := *s.getHead()
	if s.direction == right{
		newHead.x++
	} else if s.direction == left{
		newHead.x--
	} else if s.direction == up{
		newHead.y--
	} else{
		newHead.y++
	}
	if s.length > len(s.body){
		s.body = append(s.body, newHead)
	} else{
		s.body = append(s.body[1:], newHead)
	}

	s.SetPosition(newHead.x, newHead.y)
	if s.collideTest(){
		os.Exit(1)
	}
	
	for _, j:=range s.body{
		screen.RenderCell(j.x, j.y, &tl.Cell{
			Fg: tl.ColorBlack,
			Ch: '@',
		})
	}
}

func (s *Snake) Tick(event tl.Event){
	if event.Type == tl.EventKey{
		switch event.Key {
		case tl.KeyArrowRight:
			if s.direction != left {
				s.direction = right
			}
		case tl.KeyArrowLeft:
			if s.direction != right {
				s.direction = left
			}
		case tl.KeyArrowUp:
			if s.direction != down {
				s.direction = up
			}
		case tl.KeyArrowDown:
			if s.direction != up {
				s.direction = down
			}
		}
	}
}

var game *tl.Game
var field *Field

func main(){

	rand.Seed(time.Now().UnixNano())

	game := tl.NewGame()

	main := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})
	width := 80
	height := 30
	//termbox.Init()
	//width, height = termbox.Size()
	field := newField(width, height)
	snake := newSnake(field)
	food := newFood()

	main.AddEntity(field)
	main.AddEntity(snake)
	main.AddEntity(food)

	game.Screen().SetLevel(main)
	game.Screen().SetFps(10)
	game.Start()
}
