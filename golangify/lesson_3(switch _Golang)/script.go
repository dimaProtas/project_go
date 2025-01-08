package main

import (
	"fmt"
	"time"
)

// Находим день недели
func task_1() {
	switch time.Now().Weekday() {

	case time.Monday:
		fmt.Println("Сегодня понедельник!")

	case time.Tuesday:
		fmt.Println("Сегодня вторник!")

	case time.Wednesday:
		fmt.Println("Сегодня среда!")

	case time.Thursday:
		fmt.Println("Сегодня четверг!")

	case time.Friday:
		fmt.Println("Сегодня пятница!")

	case time.Saturday:
		fmt.Println("Сегодня суббота!")

	case time.Sunday:
		fmt.Println("Сегодня воскресенье!")
	}
}

// В go возможно использовать в switch несколько значений через запятую!
func task_2() {
	switch time.Now().Weekday() {
	case time.Monday, time.Wednesday, time.Friday, time.Tuesday, time.Thursday:
		fmt.Println("Сегоня будний день!")
	case time.Sunday, time.Saturday:
		fmt.Println("Сегодня выходной!")
	}
}

// Оператор switch пример с default в Go.
func task_3() {
	size := "XXS"

	switch size {
	case "XXS":
		fmt.Println("очень очень маленький размер!")
	case "XS":
		fmt.Println("очень маленький размер!")
	case "S":
		fmt.Println("маленький размер!")
	case "M":
		fmt.Println("средний размер!")
	case "L":
		fmt.Println("большой размер!")
	case "XL":
		fmt.Println("очень большой размер!")
	case "XXL":
		fmt.Println("огромный размер!")
	default:
		fmt.Println("неизвестный размер!")
	}
}

// Необязательный оператор при работе со switch в Go.
func task_4() {
	var num_input int
	fmt.Println("Ввелите число: ")
	fmt.Scan(&num_input)
	switch num := num_input; num%2 == 0 {

	case true:
		fmt.Println("Число четное!")
	case false:
		fmt.Println("Число нечетное!")
	}
}

// Оператор break в Go
func task_5() {
	w := "a b c\td\nefg hi"

	for _, e := range w {
		switch e {
		case ' ', '\n', '\t':
			break
		default:
			// fmt.Printf("%c_%c\n", e, e) // Printf не переносит на новую строку в отличии от Println
			// fmt.Printf("%c_%c,", e, e) // %c предназначен для вывода rune как символа
			fmt.Println(e)
		}
	}
}

// Оператор switch без выражения в Go.
// При использовании оператора switch без выражения, это фактически равно выражению switch true. Эту форму можно использовать вместо многострочных операторов if/else для сокращения кода.
func task_6() {
	now := time.Now()

	switch {
	case now.Hour() < 12:
		fmt.Println("AM", now.Format("2006-01-02 15:04"))

	default:
		fmt.Println("PM", now.Format("2006-01-02 15:04"))
	}
}

// Ключевое слово fallthrough в switch.
// Для перехода к следующему кейсу можно использовать ключевое слово fallthrough.
func task_7() {
	nextstop := "B"
	fmt.Println("Остановки впереди:")

	switch nextstop {
	case "A":
		fmt.Println("очтановка A")
		fallthrough
	case "B":
		fmt.Println("очтановка B")
		fallthrough
	case "C":
		fmt.Println("очтановка C")
		fallthrough
	case "D":
		fmt.Println("очтановка D")
	}
}

// Использование type в switch
func task_8() {
	var data interface{}

	data = true

	switch mytype := data.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	default:
		fmt.Printf("%T", mytype)
	}
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
