package main

import "github.com/nsf/termbox-go"

type Obstacle struct {
	x, y      int
	direction int // 0 for vertical, 1 for horizontal
}

func NewObstacle(x, y, direction int) *Obstacle {
	return &Obstacle{x: x, y: y, direction: direction}
}

func (o *Obstacle) Move() {
	width, height := termbox.Size()

	if o.direction == 0 {
		o.y++
		if o.y >= height {
			o.y = 0
		}
	} else {
		o.x--
		if o.x < 0 {
			o.x = width - 1
		}
	}
}
