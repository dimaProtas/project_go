package main

import (
	"fmt"
	"math"
)

func main() {
	println("Решите квадратное уровнение!")

	var a float64
	println("Введите число а: ")
	fmt.Scan(&a)

	var b float64
	println("Введите число b: ")
	fmt.Scan(&b)

	var c float64
	println("Введите число с: ")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Дискрименант равен" + fmt.Sprint(D) + "\nРезультат равен " + fmt.Sprint(x1) + " И " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Println("D==" + fmt.Sprint(D) + "\nКорень равен " + fmt.Sprint(x))
	} else {
		println("Уровнение не имеет корней!")
	}

	// fmt.Println("Вы ввели число a==" + fmt.Sprint(a))
	// fmt.Println("Вы ввели число b==" + fmt.Sprint(b))
	// fmt.Println("Вы ввели число c==" + fmt.Sprint(c))
}
