package main

import (
	"fmt"
	"strings"
	"time"
)

// Тип interface в Golang
//Методы выражают поведение предоставленного типа, поэтому интерфейсы объявляются с набором методов, которых тип должен удовлетворить.
// В следующем примере объявляется переменная с типом interface:
// var t interface {
// 	talks() string
// }

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "I am a martian"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("toot ", int(l))
}

// Тип интерфейса можно использовать везде, где используются другие типы. К примеру, следующая функция shout обладает параметром типа talker.
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

// Аргумент, который вы передаете функции shout, должен удовлетворять требования интерфейса talker.
// К примеру, тип crater не удовлетворяет интерфейс talker, поэтому компилятор Go откажется компилировать программу:
type crate struct{}

func (c crate) talk() string {
	fmt.Println("Теперь crate есть метод intarface, и он будет работать, interface сам подтянет метод к нужному типу, но если метода не будет, то вызвать shout не получиться!")
	return "Good"
}

// Интерфейсы показывают свою гибкость, когда вам нужно изменить или расширить код.
// При объявлении нового типа с методом talk функция shout будет с ним работать. Любой код, что зависит только от интерфейса,
// может оставаться прежним, даже если имплементации добавляются или меняются.

type starship struct {
	laser
}

// Расширьте Листинг 4, 5 и 6, объявив новый тип rover (структура) c методом talk, что возвращает «whir whir». Используйте функцию shout с вашим новым типом.
type rover struct{}

func (r rover) talk() string {
	return "whir whir"
}

func task_1() {
	// t = martian{}
	// fmt.Println(t.talks())

	// t = laser(5)
	// fmt.Println(t.talks())

	shout(martian{})
	shout(laser(4))
	shout(crate{})
	shout(rover{})

	s := starship{laser(5)}
	shout(s)

}

// Примеры использования интерфейсов в коде на Golang
type startdater interface {
	YearDay() int
	Hour() int
}

func startdate(t startdater) float64 {
	day := float64(t.YearDay())
	hoyr := float64(t.Hour()) / 24
	return 1000 + day + hoyr
}

// Функция stardate оперирует как земными датами, так и марсианскими днями, как показано в следующем листинге.
type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

func task_2() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("Выдуманная дата старта: %.2f\n", startdate(day))

	s := sol(1422)
	fmt.Printf("Интерфейс позволяет оперировать с марсианским временем блогодоря методам интерфейса: %v\n", startdate(s))

}

// Удовлетворение требований интерфейса в Go
type location struct {
	lat, long coordinate
}

func (l location) String() string {
	return fmt.Sprintf("Переопределение: %v, %v", l.lat, l.long)
}

// Напишите метод String для типа coordinate и location и используйте его для отображения координат в более читабельном формате.
type coordinate struct {
	d, m, s float64
	h       rune
}

type Stringer interface {
	String() string
}

func (c coordinate) String() string {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	res := sign * (c.d + c.m/60 + c.s/3600)
	return fmt.Sprintf("%v", res)
}

func task_3() {
	curiosiity := location{
		lat:  coordinate{4, 35, 22.2, 'S'},
		long: coordinate{137, 26, 30.12, 'E'},
	}
	fmt.Println(curiosiity)

	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	opportunity := location{
		lat:  coordinate{4, 35, 22.2, 'S'},
		long: coordinate{137, 26, 30.12, 'E'},
	}
	fmt.Println(opportunity, lat, long)

}

func main() {
	// task_1() // Тип interface в Golang
	// task_2() // Примеры использования интерфейсов в коде на Golang
	task_3() // Удовлетворение требований интерфейса в Go
}
