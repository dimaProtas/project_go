package main

import (
	"fmt"
	"sort"
	"strings"
)

// Срез массива в Golang
func task_1() {
	planets := [...]string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}

	terrastrial := planets[0:4]
	gasGigants := planets[4:6]
	iceGigants := planets[6:8]
	fmt.Println("Земная группа    Газовые гиганты    Ледяные гиганты")
	fmt.Println("===================================================")
	for i := 0; i < len(terrastrial); i++ {
		if i < 2 {
			fmt.Printf("|%13v | %13v | %10v|\n", terrastrial[i], gasGigants[i], iceGigants[i])
		} else {
			fmt.Printf("|%13v |\n", terrastrial[i])
		}

	}
	// Присваивание нового значения к элементу среза меняет базовый массив planets. Изменение будет видно и в других срезах:
	fmt.Println("===================================================")
	fmt.Println("----------- Разрезаем срез еще раз! --------------")
	terrastrialMarkII := terrastrial[2:4]
	terrastrial[2] += "-Родная"
	fmt.Printf("%15v\n", terrastrialMarkII)
	fmt.Println(planets)
	fmt.Println("--------------------------------------------------")
}

// Индексы для среза по умолчанию
func task_2() {
	planets := [...]string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}

	terrasttrial := planets[:4]
	gasGigants := planets[4:6]
	iceGigants := planets[6:]
	fmt.Printf("%v\n%v\n%v\n", terrasttrial, gasGigants, iceGigants)

	// К чему приводит пропуск обоих индексов? Переменная allPlanets является срезом, который теперь содержит все восемь планет:
	allPlanets := planets[:]
	fmt.Println(allPlanets)
}

// Срез строк в Golang
func task_3() {
	// Синтаксис среза массива также можно использовать для строк
	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println(tune)

	// Результатом среза строки является другая строка. Однако, присваивание нового значения neptune не изменит значения tune и наоборо
	neptune = "Poseidon"
	fmt.Println(tune)

	// Имейте в виду, что индексы учитывают количество байтов, но не рун
	question := "¿Cómo estás?"
	fmt.Println(question[:6])

}

// Если бы Земля и Марс были единственными колонизированными планетами, как бы вы могли получить срез colonized от terrestrial?
func task_4() {
	planets := [...]string{"Меркурий", "Венера", "Земля", "Марс", "��питер", "Сатурн", "Уран", "Нептун"}
	terrastrial := planets[:4]
	colanized := terrastrial[2:]
	fmt.Println(colanized)
}

// Композитные литералы для срезов
func task_5() {
	// Многие функции Go лучше оперируют со срезами, чем с массивами. Если вам нужен срез, что показывает каждый элемент базового массива,
	// можно объявить массив, а затем сделать срез через [:]. Это делается следующим образом:
	drawArray := [...]string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	drawSlise := drawArray[:]
	fmt.Printf("%v - Создали срез масива!\n", drawSlise)

	// В следующем примере срез dwarfs инициализируется через уже знакомый синтаксис композитного литерала:

	drawNewArray := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	fmt.Printf("%v - Обьявил сразу срез!\n", drawNewArray)
	// Базовый массив по-прежнему существует. Незаметно, Go сам объявляет массив из пяти элементов, а затем делает срез, что показывает все элементы.

	// Используйте специальный символ %T для сравнения типов dwarfArray и среза dwarfs.
	fmt.Printf("%T - тип масива dwarArray\n%T - тип среза drawNewArray\n", drawArray, drawNewArray)
}

// Преимущества среза массива в Golang
func hyperspase(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func task_6() {

	planets := []string{"      Меркурий     ", "  Венера  ", "  Земля    ", "    Марс      ", "          Юпитер        ", "      Сатурн        ", "       Уран        ", "     Нептун      "}
	hyperspase(planets)                     // Функция удаляющая пробелы у каждого  элемента среза
	worldsAdd := strings.Join(planets, ",") // Обьеденение среза в строку
	fmt.Println(worldsAdd)
	fmt.Println(strings.ReplaceAll(worldsAdd, ",", "-")) // Замена симолов в строке

	// Срезы более изменчивые, нежели массивы и в других аспектах. У срезов есть длина, однако, в отличие от длины массивов,
	// она не является частью типа. Вы можете передать срез любого размера функции hyperspace:
	drawArray := []string{"Церера", "Плутон"}
	hyperspase(drawArray)
	fmt.Printf("Передаем в функцию срез с любым количеством элементов! - %v\n", drawArray)

}

// Срезы с методами в Golang
func task_7() {
	planets := []string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}

	// Пакет sort стандартной библиотеки объявляет тип StringSlice
	// Для упорядочивания планет в алфавитном порядке в следующем примере planets конвертируется в тип sort.StringSlice, а затем вызывается метод Sort:
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)

	// В пакете sort есть вспомогательная функция Strings для конвертации типа и вызова метода Sort. Это значительно облегчает процесс:
	sort.Strings(planets)
	fmt.Printf("Короткая запись sort - %v\n", planets)

	// Сортировка чисел
	num := []int{1, 2, 3}
	sort.Ints(num)
	fmt.Println(num)
}

// Напишите программу для преобразования слайса строки через добавление слова "Новый " перед названием планеты.
// Используйте программу для изменения названий планет Марс, Уран и Нептун.
// В первой итерации может использоваться функция terraform, но в конечной реализации должен быть введен тип Planets с методом terraform,
// похожим на sort.StringSlice.

type Planets []string

func (p Planets) NewPlanets() {
	for i := range p {
		p[i] = "Новый " + p[i]
	}
}

func task_8() {
	planets := []string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}
	Planets(planets[2:4]).NewPlanets()
	Planets(planets[6:]).NewPlanets()

	// Принт выводит полный срез planets, но изменены были только те элементы которые были у казаны в срезе метода NewPlanets.
	fmt.Println(planets)
}

func main() {
	// task_1() // Срез массива в Golang
	// task_2() // Индексы для среза по умолчанию
	// task_3() // Срез строк в Golang
	// task_4() // DZ
	// task_5() // Композитные литералы для срезов
	// task_6() // Преимущества среза массива в Golang
	// task_7() // Срезы с методами в Golang
	task_8() // Напишите программу для преобразования слайса строки через добавление слова "Новый " перед названием планеты.
}
