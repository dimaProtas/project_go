package main

import "fmt"

func main() {
	// Угадай число на go
	println("Загадай число от 1 до 1000, а компьютер отгадает!")
	var num int
	fmt.Scanln(&num)

	var nums [1000]int
	var response string

	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}

	low, high := 0, len(nums)-1

	for {
		mid := (low + high) / 2
		result := nums[mid]
		fmt.Printf("Ваше число: %d?\nДайте ответ =(равно) или <(число меньше) или >(число больше) \n", result)
		fmt.Scanln(&response)

		if response == "=" {
			fmt.Printf("Ваше число: %d\n", result)
			break
		} else if response == "<" {
			println(mid)
			high = mid - 1
			println(mid)
		} else if response == ">" {
			println(mid)
			low = mid + 1
			println(mid)
		} else {
			fmt.Println("Неверный ответ!")
		}
	}
}
