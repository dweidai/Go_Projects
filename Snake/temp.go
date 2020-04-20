package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
	"github.com/nsf/termbox-go"
)

// Coord represents a point on the screen. x grows right and y grows down.
type Coord struct {
	x, y int
}
// Food handles its own collisions with a Snake and places itself randomly
// on the screen. Since there's no top-level controller, it also updates the
// score.
type Food struct {
	*tl.Entity
	coord Coord
}

// NewFood creates a new Food at a random position.
func NewFood() *Food {
	f := new(Food)
	f.Entity = tl.NewEntity(1, 1, 1, 1)
	f.moveToRandomPosition()
	return f
}

// Draw draws the Food as a default character.
func (f *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(f.coord.x, f.coord.y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '@',
	})
}

// Position returns the x,y position of this Food.
func (f Food) Position() (int, int) {
	return f.coord.x, f.coord.y
}

// Size returns the size of this Food - always 1x1.
func (f Food) Size() (int, int) {
	return 1, 1
}

// Collide handles collisions with the Snake. It updates the score and places
// the Food randomly on the screen again.
func (f *Food) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Snake:
		// It better be a snake that we're colliding with...
		f.handleSnakeCollision()
	}
}

func (f *Food) moveToRandomPosition() {
	newX := randInRange(1, border.width-1)
	newY := randInRange(1, border.height-1)
	f.coord.x, f.coord.y = newX, newY
	f.SetPosition(newX, newY)
}

func (f *Food) handleSnakeCollision() {
	f.moveToRandomPosition()
	IncreaseScore(5)
}

func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}
// Border is the edge of the playing area. If the Snake collides with it,
// it dies.
type Border struct {
	*tl.Entity
	width, height int
	coords        map[Coord]int
}

// NewBorder creates a Border with the given dimensions.
func NewBorder(width, height int) *Border {
	b := new(Border)
	b.Entity = tl.NewEntity(1, 1, 1, 1)
	// Subtract one to account for bottom and right border
	b.width, b.height = width-1, height-1

	b.coords = make(map[Coord]int)

	// Top and bottom
	for x := 0; x < b.width; x++ {
		b.coords[Coord{x, 0}] = 1
		b.coords[Coord{x, b.height}] = 1
	}

	// Left and right
	for y := 0; y < b.height+1; y++ {
		b.coords[Coord{0, y}] = 1
		b.coords[Coord{b.width, y}] = 1
	}

	return b
}

// Contains returns true if a Coord is part of the border, else false.
// Used for collision detection.
func (b *Border) Contains(coord Coord) bool {
	_, exists := b.coords[coord]
	return exists
}

// Draw draws the border on the screen. A default color is used.
func (b *Border) Draw(screen *tl.Screen) {
	if b == nil {
		return
	}

	for c := range b.coords {
		screen.RenderCell(c.x, c.y, &tl.Cell{
			Bg: tl.ColorBlue,
		})
	}
}
type direction int

const (
	right direction = iota
	left
	up
	down
)

// Snake is the snake.
type Snake struct {
	*tl.Entity
	body      []Coord
	bodyLen   int
	direction direction
}

// NewSnake creates a new Snake with a default length and position.
func NewSnake() *Snake {
	s := new(Snake)
	s.Entity = tl.NewEntity(5, 5, 1, 1)
	s.body = []Coord{
		{3, 5},
		{4, 5},
		{5, 5}, // head
	}
	// Need to track length explicitly for the case
	// where we're actively growing
	s.bodyLen = len(s.body)
	s.direction = right
	return s
}

func (s *Snake) head() *Coord {
	return &s.body[len(s.body)-1]
}

func (s *Snake) grow(amount int) {
	s.bodyLen += amount
}

func (s *Snake) isGrowing() bool {
	return s.bodyLen > len(s.body)
}

func (s *Snake) isCollidingWithSelf() bool {
	for i := 0; i < len(s.body)-1; i++ {
		if *s.head() == s.body[i] {
			return true
		}
	}
	return false
}

func (s *Snake) isCollidingWithBorder() bool {
	return border.Contains(*s.head())
}

// Draw is called every frame so it calculates new positions and checks
// for collisions in addition to just drawing the Snake.
func (s *Snake) Draw(screen *tl.Screen) {
	// Update position based on direction
	newHead := *s.head()
	switch s.direction {
	case right:
		newHead.x++
	case left:
		newHead.x--
	case up:
		newHead.y--
	case down:
		newHead.y++
	}

	if s.isGrowing() {
		// We must be growing
		s.body = append(s.body, newHead)
	} else {
		s.body = append(s.body[1:], newHead)
	}

	s.SetPosition(newHead.x, newHead.y)

	if s.isCollidingWithSelf() || s.isCollidingWithBorder() {
		EndGame()
	}

	// Draw snake
	for _, c := range s.body {
		screen.RenderCell(c.x, c.y, &tl.Cell{
			Fg: tl.ColorGreen,
			Ch: 'o',
		})
	}
}

// Tick handles keypress events
func (s *Snake) Tick(event tl.Event) {
	// Find new direction - but you can't go
	// back from where you came.
	if event.Type == tl.EventKey {
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
		case 0:
			// Vim mode!
			switch event.Ch {
			case 'h', 'H':
				if s.direction != right {
					s.direction = left
				}
			case 'j', 'J':
				if s.direction != up {
					s.direction = down
				}
			case 'k', 'K':
				if s.direction != down {
					s.direction = up
				}
			case 'l', 'L':
				if s.direction != left {
					s.direction = right
				}
			}
		}
	}
}

// Collide is called when a collision occurs, since this Snake is a
// DynamicPhysical that can handle its own collisions. Here we check what
// we're colliding with and handle it accordingly.
func (s *Snake) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Food:
		s.handleFoodCollision()
	case *Border:
		s.handleBorderCollision()
	}
}

func (s *Snake) handleFoodCollision() {
	s.grow(5)
}

func (s *Snake) handleBorderCollision() {
	EndGame()
}

var score = 0
var game *tl.Game
var border *Border
var scoreText *tl.Text

// IncreaseScore increases the score by the given amount and updates the
// score text.
func IncreaseScore(amount int) {
	score += amount
	scoreText.SetText(fmt.Sprint(" Score: ", score))
}

// EndGame should be called when the game ends due to e.g. dying.
func EndGame() {
	endLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorRed,
	})

	game.Screen().SetLevel(endLevel)
}

func main() {
	isFullscreen := flag.Bool("fullscreen", false, "Play fullscreen!")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	game = tl.NewGame()

	mainLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	width, height := 80, 30
	if *isFullscreen {
		// Must initialize Termbox before getting the terminal size
		termbox.Init()
		width, height = termbox.Size()
	}
	border = NewBorder(width, height)

	snake := NewSnake()
	food := NewFood()
	scoreText = tl.NewText(0, 0, " Score: 0", tl.ColorBlack, tl.ColorBlue)

	mainLevel.AddEntity(border)
	mainLevel.AddEntity(snake)
	mainLevel.AddEntity(food)
	mainLevel.AddEntity(scoreText)

	game.Screen().SetLevel(mainLevel)
	game.Screen().SetFps(10)
	game.Start()
}