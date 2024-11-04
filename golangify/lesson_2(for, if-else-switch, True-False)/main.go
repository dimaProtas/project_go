package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func task_1() {
	// // Вы находитесь в черной пещере!
	fmt.Println("Вы находитесь в черной пещере!")

	var command = "выйти наружу"
	var exit = strings.Contains(command, "наружу")

	fmt.Println("Вы покидаете пещеру:", exit)

	fmt.Println("На улице очень соснечно!")

	command = "надеть очки"
	var wearShades = strings.Contains(command, "очки")

	fmt.Println("Вы надеваете очки:", wearShades)
}

// // Операторы сравнения в golang
func task_2() {
	fmt.Println("На знаке снаружи написано 'Несовершеннолетним вход запрещен'.")

	var age int
	fmt.Printf("Ваш возраст? ")
	fmt.Scan(&age)

	var adult = age >= 18

	fmt.Println("В возрасте", age, ", я совершенно летний?", adult)
}

// // Условные опаераторы в golang
func task_3() {
	var command string
	fmt.Println("Введите команду:")
	fmt.Scan(&command)

	if command == "выйти" {
		fmt.Println("Вы покидаете пещеру.")
		fmt.Println("Куда вы направитесь, в горы, пещеру или лес?")
		fmt.Scan(&command)

		if command == "горы" {
			fmt.Println("Вы покидаете пещеру и идёте в горы.")
		} else if command == "пещеру" {
			fmt.Println("Вы покидаете пещеру и идёте в новую пешеру пещеру.")
		} else if command == "лес" {
			fmt.Println("Вы покидаете пещеру и идёте в лес.")
		} else {
			fmt.Println("Вы не понимаете, что делаете, и идёте в", command, "!")
		}
	} else if command == "остаться" {
		fmt.Println("Вы остаётесь в пещере, и будете там жить вечно!")
	} else {
		fmt.Println("Вы не понимаете, что делаете.")
	}
}

// // Логические операторы || и && в Golang
func task_4() {
	var year int
	fmt.Println("Введите год, а я определю высокосный он или нет!")
	fmt.Scan(&year)

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("Высокосный год!")
	} else if year%400 != 0 {
		fmt.Println("Не высокосный год!")
	} else {
		fmt.Println("Невозможно определить!")
	}
}

// // Оператор switch в Golang.
func task_5() {
	var command string
	fmt.Println("Введите команду:")
	// Создаем новый ридер для считывания строки
	reader := bufio.NewReader(os.Stdin)
	command, _ = reader.ReadString('\n')

	// Убираем символ новой строки из строки
	command = strings.TrimSpace(command)

	switch command {
	case "выйти наружу":
		fmt.Println("Вы покидаете пещеру.")
	case "остаться внутри", "остаться":
		fmt.Println("Вы остаётесь в пещере, и будете там жить вечно!")
	case "подняться в гору", "гора":
		fmt.Println("Вы поднимаетесь в гору, и будете радоваться!")
	case "пойти в лес", "лес":
		fmt.Println("Вы пойдёте в лес!")
		fallthrough
	case "грибы":
		fmt.Println("В лесу вы сможете собирать грибы!")
	default:
		fmt.Println("Вы не понимаете, что делаете, и идете в", command)
	}
}

// Цикл for в Golang.
func task_6() {
	var count int = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second) // Задержка
		count--
	}
	fmt.Println("Запуск!")

	// Не каждый запуск проходит по плану. Реализуйте обратный отсчет, где на каждую секунду приходится шанс 1 к 100,
	// что ввиду определенных обстоятельств запуск прервется, и счетчик остановится.
	// var count int = 0

	for {
		fmt.Println(count)
		count++
		time.Sleep(time.Second) // Задержка

		if rand.Intn(100) == 0 {
			fmt.Println("Запуск отменяется.")
			break // Завершение цикла
		}

		if count == 10 {
			fmt.Println("Запуск!")
			break // Завершение цикла
		}
	}
}

// Напишите программу для угадывания числа. Заставьте компьютер выбирать случайные числа из промежутка, пока он не угадает задуманное вами число.
// Данное число нужно заранее объявить в верхней части программы. Пускай на экране отображается каждая догадка с пояснением,
// больше данное число или меньше задуманного вами варианта.
func task_check() {
	var number int
	fmt.Println("Введите число от 0 до 20, которое угадает компьютер!")
	fmt.Scan(&number)

	for {
		var randNumber int = rand.Intn(21) // Генерируем случайное число от 0 до 20 включительно
		if number != randNumber {
			if number > randNumber {
				fmt.Println("Ваше число", number, "больше randNumber", randNumber, "!")
			} else if number < randNumber {
				fmt.Println("Ваше число", number, "меньше randNumber", randNumber, "!")
			}
		} else {
			fmt.Println("Ваше число", randNumber, ", копьютер отгадал!")
			break
		}
	}

}

func main() {
	// task_1()
	// task_2()
	// task_3()
	// task_4()
	// task_5()
	// task_6()
	// task_check()
}
