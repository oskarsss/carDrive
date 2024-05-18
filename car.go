package main

import "github.com/nsf/termbox-go"

type Car struct {
	x, y int
}

func NewCar(x, y int) *Car {
	return &Car{x: x, y: y}
}

func (c *Car) HandleInput(ev termbox.Event) {
	width, height := termbox.Size() // Get the terminal width and height

	switch ev.Key {
	case termbox.KeyArrowLeft:
		if c.x > 0 {
			c.x--
		}
	case termbox.KeyArrowRight:
		if c.x < width-1 {
			c.x++
		}
	case termbox.KeyArrowUp:
		if c.y > 0 {
			c.y--
		}
	case termbox.KeyArrowDown:
		if c.y < height-1 {
			c.y++
		}
	}
}

func (c *Car) CollidesWith(o *Obstacle) bool {
	return c.x == o.x && c.y == o.y
}
