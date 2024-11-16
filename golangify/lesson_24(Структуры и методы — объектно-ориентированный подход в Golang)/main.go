package main

// Запуск с экспортироваными файлами ```go run *.go``` или ```go run main.go export.go```

import (
	"fmt"
)

// Прикрепление методов к структурам в Golang
func task_1() {
	lat := Coordinate{4, 35, 22.2, 'S'}
	long := Coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.Decimal(), long.Decimal())

}

// Функции конструктора в Golang
func task_2() {
	lat := Coordinate{4, 35, 22.2, 'S'}
	long := Coordinate{137, 26, 30.12, 'E'}
	loc := NewLocation(lat, long)
	fmt.Printf("%+v, %[1]T, %T, %T\n", loc, loc.Lat, loc.Long)
}

// Классы в Golang
func task_3() {
	mars := World{radius: 3389.5}
	// zemlya := World{radius: 6378.1}

	spirit := Location{-14.5684, 175.472636}
	opportunity := Location{-1.9462, 354.4734}

	dist := mars.Distance(spirit, opportunity)

	fmt.Printf("Растояние от spirit до opportunity равно: %.2f km\n", dist)
}

func main() {
	// task_1() // Прикрепление методов к структурам в Golang
	// task_2() // Функции конструктора в Golang
	task_3() // Классы в Golang
}
