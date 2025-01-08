package main

import (
	"fmt"
	"unicode/utf8"
)

// Объявление строковых переменных Go
func task_1() {
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`
========================================================
	strings can span
	multiple lines with the
	\n escape sequence
========================================================
`)

	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
}

// Символы, коды, rune и byte
// Для вывода символов используй %c
func task_2() {
	var pi int64 = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang) // Выводит: 960 940 969 3
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang) // Выводит: π ά ω !

	tmp := '$'
	fmt.Printf("%c\n", tmp) // Выводит: $

	var star byte = '*' // Символ *
	smail := '?'
	e := 'é'

	fmt.Printf("%c kod: %[1]v,\n%c kod: %[2]v,\n%c kod: %[3]v\n", star, smail, e, star, smail, e) // Выводит: * ? ? é * ? ?s
}

// Можно ли изменять строки в Golang?
func task_3() {
	peasce := "salam"
	peasce = "salām"
	fmt.Println(peasce)

	message := "salam"
	c := message[4]
	fmt.Printf("Получили символ 4 из переменной message: %c\n", c)
}

// Задание для проверки:
// Напишите программу для вывода каждого байта (символ ASCII) слова «shalom», по одному символу на строку.
func task_4() {
	message := "shalom"
	for i := 0; i < 6; i++ {
		fmt.Printf("Символ: %c, ASCI kode: %[1]v\n", message[i])
	}
}

// Манипуляция символами с помощью шифра Цезаря
// Вопрос для проверки:
// Каким будет результат выражения c = c — ‘a’ + ‘A’, если с является 'g' в нижнем регистре?
func task_5() {
	c := 'g'
	a := 'a'
	A := 'A'
	result := c - 'a' + 'A'
	fmt.Printf("выражение c = c — ‘a’ + ‘A’ если с является 'g' в нижнем регистре: c = %v - %v + %v\n", c, a, A)
	fmt.Printf("Результат: %v\n", result)
}

// ROT13 — современный вариант шифра Цезаря
func task_6() {
	message := "uv vagreangvbany fcnpr fgngvba"
	result := ""
	for i := 0; i < len(message); i++ {
		symbol := message[i]
		if symbol >= 'a' && symbol <= 'z' {
			symbol = symbol + 13
			if symbol > 'z' {
				symbol = symbol - 26
			}
		}

		result += string(symbol)
	}
	fmt.Println(result)
}

// Декодирование строк в руны
func task_7() {
	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes")                    // Выводит: 15 bytes
	fmt.Println(utf8.RuneCountInString(question), "runes") // Выводит: 12 runes

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size) // Выводит: First rune: ¿ 2

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}

// Вопросы для проверки:
// 1.Сколько рун в английском алфавите "abcdefghijklmnopqrstuvwxyz"? Сколько байтов?
// 2.Сколько байтов в руне '¿'?
func task_8() {
	tmp := "abcdefghijklmnopqrstuvwxyz"
	// res := 0
	// for _, c := range tmp {
	// 	fmt.Printf("%c, %[1]v\n", c)
	// 	res += int(c)
	// }
	fmt.Printf("Рун в английском алфавите: %v байтов\n", len(tmp))
	fmt.Printf("Рун в символе '¿': %v байтов\n", len("¿"))
	fmt.Printf("%v\n", len("ab¿")) // Выводит: 4
}

// Итоговое задание для проверки #1:
// Расшифруйте цитату Юлия Цезаря: L fdph, L vdz, L frqtxhuhg.
// Ваша программа должна будет сдвинуть буквы верхнего и нижнего регистра на -3. Помните, что ‘a’ становится ‘x’, ‘b’ становится ‘y’,
// а ‘c’ становится ‘z’. То же самое происходит с буквами верхнего регистра.
func task_9() {
	message := "L fdph, L vdz, L frqtxhuhg."
	result := ""
	for i := 0; i < len(message); i++ {
		symbol := message[i]
		if symbol >= 'a' && symbol <= 'z' {
			symbol -= 3
			if symbol < 'a' {
				symbol += 26
			}
		} else if symbol >= 'A' && symbol <= 'Z' {
			symbol -= 3
			if symbol < 'A' {
				symbol += 26
			}
		}
		result += string(symbol)
	}
	fmt.Println(result)
}

// Зашифруйте сообщение на испанском: “Hola Estación Espacial Internacional”  через ROT13.
// Модифицируйте Листинг 7 с использованием ключевого слова range.
// Теперь, когда вы используете ROT13 c испанским текстом, ударение над буквами сохраняется.
func task_10() {
	message := "Hola Estación Espacial Internacional"

	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
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
	// task_8()
	// task_9()
	task_10()
}
