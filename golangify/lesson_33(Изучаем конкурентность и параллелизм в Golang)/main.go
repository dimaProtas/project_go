package main

import (
	"fmt"
	"image"
	"log"
	"sync"
	"time"
) // Импортирует пакет sync

// Мьютексы в Golang
var mu sync.Mutex // Объявляет мьютекс

// Пояснения к тому, посещались ли веб страницы
// Его методы могут использоваться конкурентно из нескольких горутин
type Visted struct {
	// mu охраняет карту с посещенными страницами
	mu      sync.Mutex     // Объявление мьютекса
	visited map[string]int // Объявление карты из URL (строки) ключей к значениям integer
}

type с chan Visted

// VisitLink отслеживает, сколько раз страницы с данным URL была посещена, и возвращает верное число
func (v *Visted) VisitLink(url string) int {
	v.mu.Lock()         // Блокирует мьютекс
	defer v.mu.Unlock() // Убедитесь, что мьютекс разблокирован
	count := v.visited[url]
	count++
	v.visited[url] = count // Обновляет карту
	return count
}

func task_1() {
	// mu.Lock()         // Блокирует мьютекс
	// defer mu.Unlock() // Снимает блокировку с мьютекса перед возвращением
	// Блокировка будет действовать пока функция не вернет (return) результат.
	//#######################################################################//

	// Попробуйте изменить Листинг 3 для использования техники, изученной в начале урока о многопоточности для запуска нескольких горутин,
	// которые вызывают VisitLink с разными значениями и экспериментируют со вставкой операторов Sleep в разных местах.
	// Также попробуйте убрать вызовы Lock и Unlock, чтобы посмотреть, что произойдет.

	// Мьютекс неплох для прямого использования, он является важным инструментом при написании методов,
	// что должны быть доступными для использования из нескольких
	var wf sync.WaitGroup
	v := Visted{visited: make(map[string]int)} // Создание карты

	urls := []string{
		"https://golang.org", "https://golang.org",
		"https://mts.ru", "https://mts.ru", "https://mts.ru",
		"https://megafon.ru", "https://megafon.ru",
		"https://beeline.ru", "https://beeline.ru", "https://beeline.ru",
	}

	for _, url := range urls {
		wf.Add(1)
		go func(url string) {
			defer wf.Done()
			v.VisitLink(url)
		}(url)
	}
	wf.Wait() // Ожидание завершение горутин
	fmt.Println(v.visited)
}

// Почему нужно быть осторожным с мьютексами в Golang?
func task_2() {

}

// Долгодействующие задачи в Golang
func task_3() {
	// go worker()
	// go worker2()
	// go worker3()
	// time.Sleep(2 * time.Second) // Ожидание выполнения
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}

// worker пишется как цикл for, содержащий оператор select. Цикл повторяется пока рабочий worker жив. select ожидает чего-то интересного.
// В данном случае чем-то интересным может быть команда извне. Помните, хотя работник действует независимо,
// мы все же хотим иметь возможность контролировать его. Или это может быть событие таймера, сообщающее рабочему, что пора двигать марсоход.
func worker() {
	for {
		select {
		// Ожидание канала
		}
	}
}

func worker2() {
	n := 0
	next := time.After(time.Second) // Создаем начальный канал таймера
	for {
		select {
		case <-next: // Ожидает истечение срока таймера
			n++
			fmt.Println(n)                 // Выводит число
			next = time.After(time.Second) // Создает другой канал таймера для другого события
		}
	}
}

func worker3() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("Текущая позиция ", pos)
			next = time.After(time.Second)

		}
	}
}

// Тип RoverDriver в следующем листинге содержит канал, который будет использоваться для отправки команд работнику.
// Мы будем использовать тип command, который будет содержать отправленные команды.
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
)

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
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
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// Теперь мы можем завершить тип RoverDriver, добавив методы для управления марсоходом, как показано в Листинге 10.
// Мы объявим два метода, по одному для каждой команды. Каждый метод отправляет правильную команду на канал commandc.
// Например, если мы вызываем метод Left, он отправит значение команды left, которое работник получит и изменит направление.
func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func main() {
	// task_1() // Мьютексы в Golang
	// task_2() // Почему нужно быть осторожным с мьютексами в Golang?
	task_3() // Долгодействующие задачи в Golang
}
