package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Game struct {
	car       *Car
	obstacles []*Obstacle
	running   bool
}

func NewGame() *Game {
	car := NewCar(10, 18)
	obstacles := []*Obstacle{}
	return &Game{car: car, obstacles: obstacles, running: true}
}

func (g *Game) Start() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	go g.handleInput()

	rand.Seed(time.Now().UnixNano())
	tick := time.NewTicker(100 * time.Millisecond)
	obstacleTicker := time.NewTicker(1 * time.Second)

	for g.running {
		select {
		case <-tick.C:
			g.update()
			g.render()
		case <-obstacleTicker.C:
			g.addObstacle()
		}
	}
}

func (g *Game) handleInput() {
	for g.running {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				g.running = false
			}
			g.car.HandleInput(ev)
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func (g *Game) update() {
	for _, obstacle := range g.obstacles {
		obstacle.Move()
	}
	g.checkCollisions()
}

func (g *Game) checkCollisions() {
	for _, obstacle := range g.obstacles {
		if g.car.CollidesWith(obstacle) {
			g.running = false
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()
			termbox.Close()
			println("Game Over!")
		}
	}
}

func (g *Game) addObstacle() {
	width, height := termbox.Size()
	direction := rand.Intn(2) // 0 for vertical, 1 for horizontal

	var newObstacle *Obstacle
	if direction == 0 {
		newObstacle = NewObstacle(rand.Intn(width), 0, direction)
	} else {
		newObstacle = NewObstacle(width-1, rand.Intn(height), direction)
	}

	g.obstacles = append(g.obstacles, newObstacle)
}
