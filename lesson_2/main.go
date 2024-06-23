package main

import "fmt"

func main() {
	var naame string
	fmt.Println("Как тебя зовут?")
	fmt.Scan(&naame)                // вводим имя
	fmt.Println("Привет, " + naame) // выводим имя
	fmt.Println("Сколько тебе лет?")
	var age uint16
	fmt.Scan(&age)                                  // вводим возраст
	fmt.Println("Тебе " + fmt.Sprint(age) + " лет") // выводим возраст
}
