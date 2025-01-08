package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

// /// task_1 /////
// Используя Листинг 5 в качестве стартовой точки, измените код таким образом, чтобы время отсрочки становилось на половину секунды длиннее с каждым шагом.
func dz1() {
	worker()
}

func worker() {
	pos := image.Point{X: 10, Y: 10}     // Текущая позиция (изначально [10, 10])
	direction := image.Point{X: 1, Y: 0} // Текущее направление (изначально [1, 0])
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("текущая позиция ", pos) // Выводит текущую позицию
			next = time.After(time.Second / 2)
		}
	}
}

///////////////////////////////////

// /// task_2 /////
// Используя тип RoverDriver в качестве стартовой точки, определите методы Start и Stop и ассоциируемые команды, а также заставьте марсоход подчиняться им.
func dz2() {
	r := NewRoverDriver()
	time.Sleep(1 * time.Second)
	r.Stop()
	time.Sleep(2 * time.Second)

}

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
	// dz1()
	dz2()
}
