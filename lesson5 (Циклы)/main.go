package main

import "fmt"

func main() {
	// Выводим 5 чисел
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("Число: ", i)
	// }

	// Выводим только чётные числа
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// Выводим только чётные числа вариант 2
	// for i := 1; i <= 100; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// Использование break
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		if i == 52 {
	// 			break
	// 		}
	// 		fmt.Println("Выводим: ", i)
	// 	}
	// }

	// Подобие цикла While
	// i := 0
	// for i <= 5 {
	// 	fmt.Println("Число:  ", i)
	// 	i++
	// }

	// Перебор маcива циклом for
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(len(nums))

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// Перебор маcива циклом for используя range и тут еще форматирование строки
	// nums := []string{"lox", "alex", "ban", "Miron"}
	// fmt.Println(len(nums))

	// for index, el := range nums {
	// 	fmt.Printf("Индекс: %d, Элемент %s\n", index, el)
	// }

	// Если нет необходимости использовать index стави _
	// nums := []string{"lox", "alex", "ban", "Miron"}
	// fmt.Println(len(nums))

	// for _, el := range nums {
	// 	fmt.Printf("Элемент %s\n", el)
	// }

	// Печатуем матрицу из двухмерного масива
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(len(matrix))

	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

}
