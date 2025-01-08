package main

import (
	"image"
	"log"
	"time"
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

type command int

const (
	right = command(0)
	left  = command(1)
	stop  = command(2)
	start = command(3)
)

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	speed := 1
	for {
		select {
		case c := <-r.commandc: // Отправляем команду
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case stop:
				speed = 0
			case start:
				speed = 1
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			// Метод Mul обычно умножает вектор (или другую структуру, представляющую направление) на скаляр (число).
			// В данном случае direction.Mul(speed) выполняет масштабирование вектора direction на значение speed.
			pos = pos.Add(direction.Mul(speed))
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

//////////////////////////////////////////////

func main() {
	r := NewRoverDriver()
	time.Sleep(1 * time.Second)
	r.Stop()
	time.Sleep(2 * time.Second)
}
