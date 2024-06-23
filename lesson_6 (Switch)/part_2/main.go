package main

import "fmt"

func main() {
	number := 10

	switch { // не пишем по какой переменной сравниваем, для того что бы не определять типы в case (так тоже можно)

	case number > 1:
		fmt.Println("Число больше 1")
		fallthrough // Используеться что бы при совпалении case выполнять проверку по другим case (иначе будет brake)
	case number < 11:
		fmt.Println("Число меньше  11")
	}
}
