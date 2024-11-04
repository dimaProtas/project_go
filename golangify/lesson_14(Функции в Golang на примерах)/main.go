package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Объявление функции в Golang
func task_1() {
	random := rand.Intn(100)
	fmt.Println("Random - ", random, "\nSumNum: ", sumNum(random, 2))

	timeUnix := time.Unix(1729935354, 0)
	fmt.Printf("timeUnix: %v - %[1]T\n", timeUnix)

	str := strings.Contains("Strings my bmw motocikle", "bmw")
	fmt.Println("фугкция Contains: ", str)

	//Atoi конвертирует строку в число и возвращает два результата
	num, err := strconv.Atoi("10")
	fmt.Println("Atoi: ", num, " - ", err)

	//func Println(a ...interface{}) (n int, err error)
	// Функция Println принимает один параметр а, но вы уже видели, что передача нескольких аргументов также возможна. Кроме того,
	// вы можете передать функции Println переменное количество аргументов, на данную особенность указывает многоточие (...).
	// Для этого есть специальный термин — вариативная функция, которой является Println. Параметр а является набором аргументов,
	// передаваемых функции. О вариативных функциях более детально поговорим в одном из следующих уроков.
}

func sumNum(n int, b int) int {
	res := n + b
	return res
}

// Создание функции в Golang.
func kelvinToCelsiusTask_2(k float64) (float64, string) {
	k -= 273.15
	str := "Hello world"
	return k, str
}

func task_2() {
	kelvin := 294.0
	kelvin2 := 233.0
	celsius, str := kelvinToCelsiusTask_2(kelvin)
	celsius2, str2 := kelvinToCelsiusTask_2(kelvin2)
	fmt.Printf("Kelvin = %v, Celsius = %-2.4v °C, sring = %v\n", kelvin, celsius, str)
	fmt.Printf("Kelvin = %v, Celsius = %-2.4v °C, sring = %v\n", kelvin2, celsius2, str2)
}

// Итоговое задание для проверки:
// Используйте Go Playground и модифицируйте Листинг 1 для добавления дополнительных функций конвертирования температуры:

// 1) Повторно используйте функцию kelvinToCelsius для конвертации 233° К в градусы Цельсия;
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// 2) Напишите и используйте функцию конвертации температуры в градусы Фаренгейта — celsiusToFahrenheit.
// Формула для конвертации температуры в градусы по Фаренгейту: (c * 9.0 / 5.0) + 32.0;
func celsiusToFahrenheit(c float64) float64 {
	celsius := (c * 9.0 / 5.0) + 32.0
	return celsius
}

// 3) Напишите функцию kelvinToFahrenheit и проверьте, чтобы она конвертирова 0° К в приблизительно –459.67° F.
func kelvinToFahrenheit(k float64) float64 {
	f := celsiusToFahrenheit(kelvinToCelsius(k))
	return f
}

func task_3() {
	// task_1
	celsius := 27.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("Celsius = %v, Fahrenheit = %v\n", celsius, fahrenheit)

	// task_2
	fahrenheit2 := 0
	kelvin := kelvinToFahrenheit(float64(fahrenheit2))
	fmt.Printf("Kelvin = %v, Fahrenheit = %v\n", kelvin, fahrenheit2)
}

func main() {
	// task_1()
	// task_2()
	task_3()
}
