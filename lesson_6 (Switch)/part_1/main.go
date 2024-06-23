package main

import "fmt"

func main() {
	println("Введи имя: ")
	var name string
	fmt.Scanln(&name)

	switch name {

	// case считывает переменную name и в зависимости от её значения возвращает результат
	// так же можно задать дефолтный ответ если не одно значение не совпадает!
	case "Вася":
		fmt.Println("Василий у нас мистер Пупкин!")
	case "Петя":
		fmt.Println("Петя у нас Петро!")
	case "Маша":
		fmt.Println("Маша растеряша!")
	case "Дмитрий":
		fmt.Println("Дмитрий это я!)")
	default:
		fmt.Println("Таких я не знаю!")
	}
}
