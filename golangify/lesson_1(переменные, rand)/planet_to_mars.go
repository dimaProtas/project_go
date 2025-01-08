package main

import (
	"fmt"
)

func main() {
	// Мой вариант
	const distance, day = 56000000, 28
	var lightSpeed int
	var hoursPer int = day * 24

	lightSpeed = distance / hoursPer

	fmt.Println(lightSpeed, "км/ч.")

	// Вариант с урока
	// const hoursPerDay = 24

	// var days = 28
	// var distance = 56000000 // km

	// fmt.Println(distance/(days*hoursPerDay), "км/ч")

}
