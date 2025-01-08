package main

import (
	"fmt"
)

func main() {
	// Весс на Марсе
	// var name string
	// var mass float64
	// var age float64

	// fmt.Println("Как тебя зовут?")
	// fmt.Scan(&name)

	// fmt.Print("Какой твой вес?\n")
	// fmt.Scan(&mass)

	// fmt.Print("Твой возраст?\n")
	// fmt.Scan(&age)

	// mass *= 0.3783
	// age = age * 365 / 687

	// fmt.Println(fmt.Sprint(name) + " твой вес на поверхности Марса равен " +
	// 	fmt.Sprint(mass) + "кг, а возраст равен " + fmt.Sprint(age) + " годам. ")

	// Использование Printf и отступ $%4v слева, отступ $%-4v\n справа
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	// Использование переменных
	// const hoursPerDay, minutsPerHour = 24, 60
	// const lightSpeed = 100800 // км/с
	// var distance = 96300000   // км

	// fmt.Println(((distance / lightSpeed) / hoursPerDay), "дней")

	// fmt.Println(((distance / lightSpeed) * minutsPerHour), "минут")

	// distance = 401000000
	// fmt.Println(distance/lightSpeed, "секунд") // В результате 1337 секунд

	// Напишите кратчайшую строку кода для вычитания двух фунтов из переменной под названием weight.
	// var weight float64 = 200
	// weight -= 2
	// fmt.Println(weight, " фунтов")

	// Случайное число в go модуль math/rand
	// var num = rand.Intn(10) + 1
	// fmt.Println(num)

	// num = rand.Intn(10) + 1
	// fmt.Println(num)

	// Расстояние между Землей и Марсом в разное время отличается и зависит от того,
	// где планеты в данный конкретный момент времени находятся на орбите Солнца.
	// Напишите программу для генерации случайного расстояния в промежутке от 56 000 000 до 401 000 000 км.
	// var distance = rand.Intn(345000001) + 56000000
	// fmt.Println(distance)

}
