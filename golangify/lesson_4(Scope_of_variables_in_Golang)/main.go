package main

import (
	"fmt"
	"math/rand"
	"time"
)

// // Область видимости переменных в Golang
func task_1() {
	count := 0

	for count < 10 {
		num := rand.Intn(10) + 1
		fmt.Println(num)

		count++
	}
}

// пример 2
func task_2() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Zero")
	} else if num == 1 {
		fmt.Println("One")
	} else {
		fmt.Println("Two")
	}
}

// Локальная и глобальная область видимости (тут с датoй поработал)
func task_3() {
	for count := 0; count < 10; count++ {
		year := rand.Intn(9) + 2020

		var era string = "AD"
		var daysIsMoonth int = 31
		moonth := rand.Intn(12) + 1
		switch moonth {
		// case 2:
		// 	daysIsMoonth = 28
		case 4, 6, 9, 11:
			daysIsMoonth = 30
		}

		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			daysIsMoonth = 29
			fmt.Println(year, "Высокосный")
		} else {
			daysIsMoonth = 28
			fmt.Println(year, "Не высокосный")
		}

		day := rand.Intn(daysIsMoonth) + 1
		date := time.Date(year, time.Month(moonth), day, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC).Format("2006-01-02 15:04:05")

		fmt.Printf("era: %v, %v\n", era, date)
	}

}

func main() {
	// task_1()
	// task_2()
	task_3()
}
