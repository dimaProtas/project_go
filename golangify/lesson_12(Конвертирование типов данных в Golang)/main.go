package main

import (
	"fmt"
	"math"
	"strconv"
)

func task_1() {
	num := 10
	number := string(num)
	str := "Hellow " + "Dima" + "сейчас " + number + " число!"
	fmt.Println(str)

	age := 41
	marsAge := float64(age)
	result := age > int(marsAge)
	fmt.Println(result)

	red := 8.0
	red_new := uint8(red)
	fmt.Printf("%v, %[1]T\n", red_new)

	// num := 10
	new_srr := strconv.Itoa(num)
	fmt.Printf("%v, %[1]T\n", new_srr)
}

func task_2() {
	var bh float64 = 32767
	var h = int16(bh) // TODO: добавить формулы ракетостроения

	if bh < math.MinInt16 || bh > math.MaxInt16 {
		fmt.Println("Переполнение диапазона!")
		// решает проблему значения за пределами допустимого диапазона
	} else {
		fmt.Println(h)

	}
}

func task_3() {
	fmt.Println("Введите число, а я проверю находится ли оно в диапазоне uint8!")
	var v int64
	fmt.Scanln(&v)
	if v >= 0 && v <= math.MaxUint8 {
		fmt.Printf("Диапазон в пределах uint8 , ваше число %v\n", v)
	} else {
		v8 := uint8(v)
		fmt.Println("Число", v, "Диапазон переполнен!, converte: ", v8)
	}
}

// Конвертация строк в Golang
func task_4() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v, rune: %[1]c, %v, rune: %[2]c, %v, rune: %[3]c, %v, rune: %[4]c.\n", pi, alpha, omega, bang)
	fmt.Print(string(pi), string(alpha), string(omega), string(bang), "\n") // Выводит: πάω!

	// contdown := 10
	// str := "Запуск в Т минус " + strconv.Itoa(contdown) + " секунд"
	// fmt.Println(str)

	// countdown := 9
	// str := fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	// fmt.Println(str) // Выводит: Launch in T minus 9 seconds.

	// Есть еще один способ. В пакете strconv есть функция Atoi (ASCII to integer, то есть от ASCII к целому числу).
	// Из-за того, что в строке могут быть какие-то кракозябры или слишком крупные числа, функция Atoi может вернуть ошибку:
	coutdown, err := strconv.Atoi("2343425356345asd234234")
	if err != nil {
		fmt.Println("Что то пошло не так!")
	} else {
		fmt.Printf("Число: %v\n", coutdown)
	}

}

// Конвертация булевых значений Golang.
func task_5() {
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Printf("%v, %[1]T\n", launch)
	fmt.Printf("%v, %[1]T\n", launchText)

	var yesNo string
	if launch {
		yesNo = "Да"
	} else {
		yesNo = "Нет"
	}
	fmt.Println(launch, yesNo)
}

// Обратная конвертация требует меньшее количество кода, так как вы можете присвоить результат условия напрямую переменной,
// как показано в следующем коде:
func task_6() {
	yesNo := "Нет"

	launch := (yesNo == "Да")
	fmt.Println(launch)
}

// Как можно конвертировать булев тип в целое число, чтобы 1 соответствовала значению true, а 0 — false?
func task_7() {
	uno := 0

	trueFalse := (uno == 1)
	fmt.Print(trueFalse, "\n")
}

// Напишите программу, что конвертирует строки в булевы значения:
// Строки «true», «yes» или «1» соответствуют значению true;
// Строки «false», «no» или «0» соответствуют значению false;
// Для других значений выводит сообщение об ошибке.
// Обратите внимание, что здесь можно использовать оператор switch, что принимает по несколько значений на случай case.
func task_8() {
	fmt.Println("Введлите отризательное («true», «yes» или «1»), или положительное («false», «no» или «0» ) значение!")
	var text string
	fmt.Scan(&text)

	var trueFalse bool
	switch text {
	case "true", "yes", "1":
		trueFalse = true
	case "false", "no", "0":
		trueFalse = false
	default:
		fmt.Println("Ошибка!")
	}

	fmt.Printf("строка: %v, %[1]T, converte: %v, %[2]T\n", text, trueFalse)
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
