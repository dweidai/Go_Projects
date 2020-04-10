package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)
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
	direction direciton
	length int
}

func NewSnake(width int, height int) *Snake{
	s := new(Snake)
	s.Entity = tl.newEntity(5,5,1,1)
	sX = rand.Intn(width-10)
	sY = rand.Intn(height)
	s.body = []Node{ 
		{sX, sY},
		{sX+1, sY}
		{sX+2, sY}
	}
	s.length = len(s.body)
	s.direction = right
	return s
}

func (s *Snake) getHead() *Node{
	return &s.body[length-1]
}

func (s *Snake) grow(){
	s.length ++
}

func (s *Snake) collideTest(field *Field) bool{
	if field.Contain(*s.getHead()){
		return true
	} else{
		for i:=0; i<length-1; i++{
			if *s.getHead() == s.body[i]{
				return true
			}
		}
		return false
	}
}

func (s *Snake) update(screen *tl.Screen){
	newHead := *s.getHead()
	if s.direction == right{
		newHead.x++
	} else if s.direction == left{
		newHead.x--
	} else if s.direction == up{
		newHead.y++
	} else{
		newHead.y--
	}
	if s.length > len(s.body){
		s.body = append(s.body, newHead)
	}else{
		s.body = append(s.body[1:], newHead)
	}
	if s.collideTest{
		//EndGame()
	}

	for i, j:=range s.body{
		if i!=nil{
			screen.RenderCell(c.x, c.y, &tl.Cell{
				Fg: tl.ColorGreen,
				Ch: 'o',
			})
		}
	}
}

func (s *Snake) keyPress(event tl.event){
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
