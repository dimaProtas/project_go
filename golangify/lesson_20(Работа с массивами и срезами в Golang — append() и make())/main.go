package main

import (
	"fmt"
)

// Функция append в Go
func task_1() {
	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	dwarfs = append(dwarfs, "Оркус")
	fmt.Printf("Добавляем в срез 'dwarfs' новый элемент - %v\n", dwarfs)

	// Функция append является вариативной, как и Println. Вы можете передать сразу несколько элементов для добавления:
	dwarfs = append(dwarfs, "Салация", "Квавар", "Седна")
	fmt.Printf("Добавляем в срез 'dwarfs' 3 новых элемента - %v\n", dwarfs)
	fmt.Printf("Длинна среза 'dwarfs' - %v\n", len(dwarfs))
}

// Длина и вместимость среза в Golang
// В следующем примере объявляется функция для вывода длины и вместимости среза:
func dump(lable string, slise []string) {
	fmt.Printf("%v: длинна %v, вместимость %v - %v\n", lable, len(slise), cap(slise), slise)
}
func task_2() {
	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	dump("dwarfs", dwarfs)
	dump("dwarfs", dwarfs[1:2])
}

// Разбор функции append в Go
func task_3() {
	dwarfs1 := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}

	dwarfs2 := append(dwarfs1, "Оркус")

	dwarfs3 := append(dwarfs2, "Салация", "Квавар", "Седна")
	dwarfs3[1] = "New-" + dwarfs3[1] // Если изменить третий срез то из за того что он ссылаеться на второй срез, поменяются оба

	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarsfs3", dwarfs3)

}

// Трех-индексный срез в Golang
func task_4() {
	planets := []string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}

	terrestrial := planets[0:4:4] // Фиксирует вместимость среза, и даже когда добавляем новый елемент в world который ссылаеться на terrestrial то его вместимость остаетсья 4
	world := append(terrestrial, "Церера")
	world[1] = "New-" + world[1] // Не изменяться срез terrestrial на который ссылаеться world, т.к. у него фиксированая вместимоть

	dump("terrestrial", terrestrial)
	dump("world", world)
	dump("planets", planets)
}

// Предварительное выделение срезов через make в Go
func task_5() {
	// Функция make в следующем примере уточняет длину (0) и вместимость (10) среза dwarfs. Перед заполнением dwarfs можно добавить до 10 элементов,
	// после чего append должен будет выделить новый массив.
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Церера", "Плутно", "Хаумеа", "Макемаке", "Эрида")

	dump("dwarfs", dwarfs)
}

// Объявление вариативных функций в Golang
func terraform(prefix string, world ...string) []string {
	newWorlds := make([]string, len(world))

	for i := 0; i < len(world); i++ {
		newWorlds[i] = prefix + " " + world[i]
	}
	return newWorlds
}

func task_6() {
	// Параметр worlds является срезом строк, что содержит ноль или более аргументов, передаваемых в terraform:
	terra := terraform("New", "Меркурий", "Венера", "Земля", "Марс")
	fmt.Println(terra)

	// Для передачи среза вместо множества аргументов, расширьте срез через многоточие:
	planets := []string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}
	pl := terraform("New", planets...)
	// Если бы terraform модифицировала элементы параметра worlds, срез planets также зафиксировал бы эти изменения.
	// Через использование newWorlds функция terraform избегает изменения и передачи аргументов.
	fmt.Println(pl)
	fmt.Println(planets)
}

// Итоговое задание для проверки:
// Напишите программу, что использует цикл для продолжающегося добавления элементов в срез.
// Каждый раз при изменении вместимости среза выводится новое значение. Всегда ли append удваивает вместимость при завершении места в базовом массиве?
func task_7() {
	arr := make([]int, 0, 10)
	lastCap := cap(arr)
	for i := 1; i < 1000; i++ {
		arr = append(arr, i)
		// fmt.Printf("Длинна - %v, Вместимость - %v\nМассив - %v\n", len(arr), cap(arr), arr)
		if cap(arr) != lastCap {
			fmt.Println(cap(arr), "вместительность!")
			lastCap = cap(arr)
		}
	}
}

// Решение преподователя
func task_8() {
	s := []string{}
	lastCap := cap(s)
	fmt.Println(lastCap, "Первая вместительность!")

	for i := 0; i < 10000; i++ {
		s = append(s, "An element")
		if cap(s) != lastCap {
			fmt.Println(cap(s))
			lastCap = cap(s)
		}
	}
}

func main() {
	// task_1() // Функция append в Go
	// task_2() // Длина и вместимость среза в Golang
	// task_3() // Разбор функции append в Go
	// task_4() // Трех-индексный срез в Golang
	// task_5() // Предварительное выделение срезов через make в Go
	// task_6() // Объявление вариативных функций в Golang
	task_7() // Итоговое задание для проверки
	// task_8() // Решение преподователя
}
