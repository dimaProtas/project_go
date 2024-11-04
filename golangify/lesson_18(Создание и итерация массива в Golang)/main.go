package main

import (
	"fmt"
)

// Объявление массива и получение доступа к его элементам
func task_1() {
	var planets [8]string
	// Запись елементов в масив
	planets[0] = "Меркурий"
	planets[1] = "Венера"
	planets[2] = "Земля"
	planets[3] = "Марс"

	// получение элемента из масива
	mars := planets[3]
	fmt.Println(mars)

	fmt.Println(len(planets))
	fmt.Printf("Четвертый элемент в масиве пустой? - %v\n", planets[4] == "")

}

// Диапазон значений массива в Golang
func task_2() {
	var planets [8]string = [8]string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Нептун", "Плутон"}

	// У массива из восьми элементов индексы от 0 до 7.
	// При попытке получить доступ к элементу за пределами диапазона массива компилятор Go сообщит об ошибке:
	// fmt.Println(planets[8])

	// Если компилятор Go не в состоянии зафиксировать ошибку, во время запуска программы может произойти сбой:
	i := 8
	fmt.Println(planets[i])
}

// Инициализация массивов через композитные литералы в Go
func task_3() {
	dwarfs := [5]string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	fmt.Println(dwarfs)

	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}
	fmt.Printf("Масив planets имеет длинну: %v\n", len(planets))

}

// Итерация через массивы в Go
func task_4() {
	dwarfs := [5]string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	// Итерация через каждый элемент массива напоминает итерацию каждого символа строки.
	// Мы ранее говорили об этом в уроке о строках в Golang. Это показано в примере ниже:
	for i := 0; i < len(dwarfs); i++ {
		el := dwarfs[i]
		fmt.Println(el)
	}

	// Ключевое слово range возвращает индекс и значение каждого элемента массива посредством использования меньшего количества кода и
	// меньшей вероятностью совершения ошибок, что показано в коде ниже:
	for c1, c2 := range dwarfs {
		fmt.Printf("Индекс: %v, Значение: %v\n", c1, c2)
	}

}

// Копирование массивов в Golang
func terraform(planets [8]string) [8]string {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
	return planets
}

func task_5() {
	// Присваивание массива новой переменной или передача его функции приводит к копированию всего его содержимого, что показано в следующем примере:
	planets := [...]string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Нептун", "Плутон"}

	planetsMarkII := planets

	planets[2] = "Упс"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	// Функция terraform оперирует с копией массива planets, поэтому модификации не затронут planets в функции main.
	fmt.Println(terraform(planets)) // Возвращаем изменненый масив
	fmt.Println(planets)            // Оригинальный масив не изменился

	// dwarfs := [5]string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	// terraform(dwarfs) // Нельзя использовать dwarfs (типа [5]string) как тип [8]string в качестве аргумента terraform
}

// Массивы из массивов в Golang
func task_6() {
	// Шахматная доска 8 х 8 представлена в следующем примере как массив из массива строк:
	var bard [8][8]string

	bard[0][0] = "r"
	bard[0][7] = "r"

	for cl := range bard {
		bard[1][cl] = "p"
	}

	fmt.Println(bard)
}

// Подумайте об игре Судоку. Как можно объявить сетку целых чисел размером 9 х 9?
func task_check() {
	var sudoku [9][9]int

	for i := range sudoku {
		sudoku[1][i] = 1
	}

	fmt.Println(sudoku)

	// Перебор строки циклом for range (не относиться к дз!)
	str := "STRING"
	for n, v := range str {
		fmt.Printf("Индекс: %v, Значение: %v\n", n, v)
	}
}

func main() {
	// task_1()
	// task_2()
	// task_3()
	// task_4()
	// task_5()
	// task_6()
	task_check()
}
