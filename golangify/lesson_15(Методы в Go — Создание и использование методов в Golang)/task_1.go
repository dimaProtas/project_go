package main

import (
	"fmt"
)

// Объявление новых типов Golang
func task_1() {
	type celsius float64
	var dreges celsius = 20

	var temperature celsius = dreges

	temperature += 10

	fmt.Println(temperature)
}

func task_2() {
	type celsius float64
	type fahrenheit float64

	// var c celsius = 20
	// var f fahrenheit = 20
	// Ошибка в операции: несовпадение типов celsius и fahrenheit
	// if c == f {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }
	// c += f

}

func main() {
	task_1()
	// task_2()
}
