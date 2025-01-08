package main

import (
	"fmt"
	"math/rand"
)

// Присваивание функции переменной в Go
// type kelvin float64

// func realSensor() kelvin {
// 	return 0
// }

// func fakeSensor() kelvin {
// 	return kelvin(rand.Intn(151) + 150)
// }

// func main() {
// 	// Если не полагаться на автоматическое присваивание типа, переменная сенсора может быть объявлена следующим образом:

// 	var sensor func() kelvin = fakeSensor
// 	fmt.Printf("Фейковый сенсор показывает - %v°K\n", sensor())

// 	sensor = realSensor
// 	fmt.Printf("Реальный сенсор показывает - %v°K\n", sensor())
// }

// Передача функции другой функции в Golang
// type kelvin float64

// func measureTemperature(saple int, sensor func() kelvin) {
// 	for i := 0; i < saple; i++ {
// 		k := sensor
// 		timeNow := time.Now().Format("2006-01-02 15:04:05") // формат времени и даты!
// 		fmt.Printf("Температура в %v равна - %v°K\n", timeNow, k())
// 		time.Sleep(time.Second)
// 	}
// }

// func fakeSensor() kelvin {
// 	return kelvin(rand.Intn(151) + 150)
// }

// func main() {
// 	measureTemperature(10, fakeSensor)
// }

// Объявление типов функции в Golang

// Перепишите следующую сигнатуру функции для использования типа функции: func drawTable(rows int, getRow func(row int) (string, string))
// var rows int

// type getRowMy func(row int) (string, string)

// func getRow(row int) (string, string) {
// 	res := rand.Intn(row) + 1
// 	return fmt.Sprint("Row:"), fmt.Sprintf("%v", res)
// }

// func drawTable(rows int, getRow getRowMy) string {
// 	_, res2 := getRow(rows)
// 	return fmt.Sprintf("Из количества строк - %v - получили строку № %v\n", rows, res2)
// }

// func main() {
// 	// num := 10
// 	row := drawTable(10, getRow)
// 	fmt.Println(row)
// }

// Замыкания (Closures) и анонимные функции Golang
// Вы можете присвоить анонимную функцию переменной, и затем использовать данную переменную как любую другую функцию, как показано в примере ниже:
// var f = func() {
// 	fmt.Println("Ананимная функция присвайваеться переменной!")
// }

// func main() {
// 	fmt.Println("Вызываем ананимную функцию!")
// 	f()

// 	//Объявленная переменная может быть в области видимости пакета или внутри функции, как показано в листинге ниже:
// 	var f_new = func(message string) {
// 		fmt.Println(message)
// 	}
// 	f_new("Ананимная функция может присвайваться внутри функции и там же вызываться!")

// 	//Можно объявить и задействовать анонимную функцию как показано ниже:
// 	func() {
// 		fmt.Println("Можно обьявить ананимную функцию и вызвать ее без присвайвания переменной!")
// 	}()
// }

// В Листинге 6 функция calibrate настраивается на ошибки в показаниях температуры воздуха.
// С помощью функции первого класса calibrate принимает фейковый или реальный сенсор в качестве параметра и возвращает функцию замены.
// Всякий раз, когда вызывается новая функция sensor, вызывается исходная функция, а чтение корректируется смещением.
// type kelvin float64
// type sensor func() kelvin

// func realSensor() kelvin {
// 	return 0
// }

// func calibrate(s sensor, offset kelvin) sensor {
// 	return func() kelvin {
// 		return s() + offset
// 	}
// }

// func main() {
// 	sensor := calibrate(realSensor, 10)
// 	fmt.Println(sensor())

// 	// Из-за того, что замыкание сохраняет ссылку на ближайшие переменные, а не копирует их значения,
// 	// изменения с этими переменными отражаются в вызовах к анонимной функции:
// 	var k kelvin = 294.0

// 	k++
// 	fmt.Printf("Принт переменной - k + 1 (%v)\n", k)

// 	f := func() kelvin {
// 		return k
// 	}
// 	fmt.Printf("Вызов ананимной функции - %v\n", f())
// }

// Вместо передачи 5 в качестве аргумента для calibrate, объявите и передайте переменную. Измените переменную и убедитесь,
// что результатом вызовов к sensor() по-прежнему является 5. Все из-за того, что параметр offset является копией аргумента (переданный через значение);
// type kelvin float64
// type sensor func() kelvin

// func realSensor() kelvin {
// 	return 0
// }

// func calibrate(s sensor, offset kelvin) sensor {
// 	return func() kelvin {
// 		return s() + offset
// 	}
// }

// func main() {
// 	var offset kelvin = 10
// 	sensor := calibrate(realSensor, offset)
// 	offset = 112
// 	fmt.Println(sensor())
// }

// Используйте calibrate с функцией fakeSensor из Листинга 2 для создания новой функции sensor.
// Вызовите новую функцию sensor несколько раз, обратите внимание, что оригинальная функция fakeSensor по-прежнему вызывается каждый раз,
// выдавая в результате случайные значения.
type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 10
	sensor := calibrate(fakeSensor, offset)
	offset = 112
	for i := 0; i < 10; i++ {
		fmt.Printf("Вызываем fakeSensor c каректировкой offset - %v\n", sensor())
	}
}
