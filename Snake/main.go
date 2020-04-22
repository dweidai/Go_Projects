package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
	//"github.com/nsf/termbox-go"
)

var width int
var height int

type Node struct{
	x int
	y int
}

type Food struct{
	*tl.Entity
	location Node
}

func newFood() *Food{
	f := new(Food)
	f.Entity = tl.NewEntity(1,1,1,1)
	width = 80
	height = 30
	newX := randInRange(1, width-1)
	newY := randInRange(1, height-1)
	f.location.x = newX
	f.location.y = newY
	return f
}

func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func (f *Food) Position()(int, int){
	return f.location.x, f.location.y
}

func (f *Food) Draw(screen *tl.Screen){
	screen.RenderCell(f.location.x, f.location.y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '@',
	})
}

func (f *Food) snakeCollision() {
	f.location.x = 0
	f.location.y = 0
	f = newFood()
}

func (f *Food) Collide(collision tl.Physical){
	switch collision.(type){
	case *Snake:
		f.snakeCollision()
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
				Bg: tl.ColorBlue,
		})
	}
}

func (f *Field) Contains(cord Node) bool{
	_, ok := f.edges[cord]
	return ok
}

const(
	right direction  =  iota
	left
	up 
	down 
)

type direction int

type Snake struct{
	*tl.Entity
	body []Node
	direction direction
	length int
	field *Field
}

func newSnake(field *Field) *Snake{
	s := new(Snake)
	s.Entity = tl.NewEntity(5,5,1,1)
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

func (s *Snake) grow(){
	s.length += 2
}

func (s *Snake) collideTest() bool{
	for i:=0; i<s.length-1; i++{
		if *s.getHead() == s.body[i]{
			return true
		}
	}
	if s.field.Contains(*s.getHead()){
		return true
	}
	return false
}

func (s *Snake) Collide(collision tl.Physical){
	switch collision.(type){
	case *Food:
		s.foodCollision()
	case *Field:
		s.borderCollision() 
	}
}

func (s *Snake) foodCollision(){
	s.grow()
}

func (s *Snake) borderCollision(){
	endGame()
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
		endGame()
	}
	
	for i, j:=range s.body{
		if i >= 0{
			screen.RenderCell(j.x, j.y, &tl.Cell{
				Fg: tl.ColorGreen,
				Ch: 'D',
			})
		}
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

func endGame(){
	end := tl.NewBaseLevel(tl.Cell{
		Bg:tl.ColorBlack,
	})
	game.Screen().SetLevel(end)
}

func main(){
	rand.Seed(time.Now().UnixNano())

	game := tl.NewGame()

	main := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
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
