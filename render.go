package main

import "github.com/nsf/termbox-go"

func (g *Game) render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw car
	termbox.SetCell(g.car.x, g.car.y, 'C', termbox.ColorGreen, termbox.ColorDefault)

	// Draw obstacles
	for _, obstacle := range g.obstacles {
		termbox.SetCell(obstacle.x, obstacle.y, 'O', termbox.ColorRed, termbox.ColorDefault)
	}

	termbox.Flush()
}
