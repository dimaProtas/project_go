package main

import (
	"fmt"
	"math/big"
)

// Работа с крупными числовыми значениями в Go
func task_1() {
	const lightSpeed = 299792
	const secondsPreyDay = 86400

	var distance int64 = 47.3e12

	fmt.Printf("Растояние до Альфа-центавры: %v км\n", distance)
	fmt.Println("Расстояние до Альфа Центавры составляет", distance, "км.")

	days := distance / lightSpeed / secondsPreyDay
	fmt.Printf("Полёт до Альфа-центавры со скоростью %v км/ч ззаймет %v дней \n", lightSpeed, days)

	var distance_andrameda uint64 = 2e18
	fmt.Printf("Растояние до галактики Андромеда: %v км\n", distance_andrameda)
}

// Задание для проверки:
// Расстояние от Земли до Марса варьируется от 56 000 000 км до 401 000 000 км.
// Представьте эти два значения в виде целых числе с помощью синтаксиса экспоненты (е).
func task_2() {
	var distance_mars_min uint32 = 56e6
	fmt.Printf("Минимальное растояние от земли до марса составляет: %v км\n", distance_mars_min)

	var distance_mars_max uint32 = 401e6
	fmt.Printf("Максимальное растояние от земли до марса составляет: %v км\n", distance_mars_max)
}

// Пакет big для крупных чисел в Golang
func task_3() {
	lightSpeed := big.NewInt(299792)
	secondsPreyDay := big.NewInt(86400)
	fmt.Printf("Скорость света %v км/ч, кол-во секунд в дне %v. \n", lightSpeed, secondsPreyDay)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Printf("Растояние до галактики Андрамеда: %v км\n", distance)

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPreyDay)

	fmt.Printf("Полет до галактики Андрамеда со скоростью %v км/ч ззаймет %v дней \n", lightSpeed, days)
}

// Константы нестандартного размера в Golang
func task_4() {
	const lightSpeed = 299792
	const secondsPreyDay = 86400
	const distance = 24000000000000000000

	const days = distance / lightSpeed / secondsPreyDay
	fmt.Printf("Полет до галактики Андрамеда со скоростью %v км/ч ззаймет %v дней \n", lightSpeed, days)
}

// Карликовая галактика в Большом Псе является ближайшей известной к Земле галактикой,
// что находится на расстоянии 236 000 000 000 000 000 км от нашего Солнца.
// Используйте константы для конвертации данного расстояния в световые годы.
func task_5() {
	const lightSpeed = 299792
	const secondsPreyDay = 86400
	const distance = 236000000000000000

	const secondPreyYear = secondsPreyDay * 365

	const year_to_litl_galaktic = distance / lightSpeed / secondPreyYear
	fmt.Printf("Полет до галактики Андрамеда со скоростью %v км/ч ззаймет %v световых лет \n", lightSpeed, year_to_litl_galaktic)

}

func main() {
	// task_1()
	// task_2()
	// task_3()
	// task_4()
	task_5()
}
