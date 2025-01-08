package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Тип данных
func task_1() {
	n := "2024"
	fmt.Printf("Type %T for %v\n", n, n)
}

func task_2() {
	n := 333.33
	fmt.Printf("Type %T for %v\n", n, n)
	b := "333.33"
	fmt.Printf("Type %T for %v\n", b, b)
	c := 333
	fmt.Printf("Type %T for %v\n", c, c)
	d := true
	fmt.Printf("Type %T for %v\n", d, d)
}

func task_3() {
	var red, green, blue uint8 = 0, 141, 213
	fmt.Printf("%v %v %v\n", red, green, blue)
	fmt.Printf("%x %x %x\n", red, green, blue)            // Отображение чмсел в шестнадцатеричной системе
	fmt.Printf("color: #%02x%02x%02x;", red, green, blue) // Выводит: color #008dd5;
}

// Целочисленное переполнение в Go
func task_4() {
	var red uint8 = 255
	red++
	fmt.Println(red) // Выводит: 0

	var number int8 = 127
	number++
	fmt.Println(number) // Выводит: -128
}

// Биты целочисленных значений (посмотреть биты числа)
func task_5() {
	var green uint8 = 3
	fmt.Printf("%08b\n", green) // Выводит: 00000011
	green++
	fmt.Printf("%08b\n", green) // Выводит: 00000100
}

func task_6() {
	// добавление числа больше, чем 1
	var red uint8 = 255
	red += 2
	fmt.Println(red) // Выводит: 1

	var number int8 = 127
	number += 3
	fmt.Println(number)

	// переполнение с другой стороны
	red = 0
	red--
	fmt.Println(red) // Выводит: 255

	number = -128
	number--
	fmt.Println(number)

	// переполнения 16-битного неподписанного целого числа
	var green uint16 = 65535
	green++
	fmt.Println(green) // Выводит: 0

	var blue uint8 = 255
	fmt.Printf("%08b\n", blue) // Выводит: 11111111
	blue++
	fmt.Printf("%08b\n", blue) // Выводит: 00000000

}

// Как избежать переполнения по времени в Go
func task_7() {
	future := time.Unix(12622780800, 0)
	fmt.Println(future) // Выводит: 2370-01-01 00:00:00 +0000 UTC в Go Playground
}

func task_8() {
	var payBang uint16 = 0
	for payBang < 2500 {
		pay := rand.Intn(3) + 1
		switch pay {
		case 1:
			payBang += 5
		case 2:
			payBang += 10
		case 3:
			payBang += 25
		}
		dollars := payBang / 100
		cents := payBang % 100
		fmt.Println("=========== ==========")
		fmt.Printf("Ваш баланс: %06.2d.%02d$\n", dollars, cents)
	}
	// n := 5 / 100
	// fmt.Printf("%v\n", n)
}

func main() {
	// task_1()
	// task_2()
	// task_3()
	// task_4()
	// task_5()
	// task_6()
	// task_7()
	task_8()
}
