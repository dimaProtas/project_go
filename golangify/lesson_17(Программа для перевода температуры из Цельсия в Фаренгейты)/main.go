package main

import (
	"fmt"
)

// Программа должна построить две таблицы. В первой таблице два столбца, в первом значится температура по Цельсию °C,
// а во втором — по Фаренгейту °F. Значения должны быть от 40° C до 100° C шагами в 5°.
// Для заполнения столбцов требуется использовать методы конвертации, описанные в уроке о методах.

// После заполнения одной таблицы заполните вторую таким образом, чтобы столбцы были инвертированы.
// То есть конвертация должна проводиться из градусов по Фаренгейту в градусы по Цельсию.

// Код, что вы напишите для создания таблицы, в будущем можно будет использовать вновь,
// уже для других программ, содержимое которых нужно отобразить в таблице с двумя столбцами.
// Используйте функции для разделения кода который создает таблицы от кода для вычисления значений температуры каждой строки.

// Реализуйте функцию drawTable, что принимает функцию первого класса в качестве параметра и вызывает ее для получения данных каждой созданной строки.
// Результатом передачи другой функции к drawTable должны быть другие отображаемые данные.
// мой вариант №1
// type celsius float64
// type fahrenheit float64

// type celsiusF func() celsius

// func getCelsius() celsius {
// 	return 40
// }

// func getFahrenheit(c celsius) fahrenheit {
// 	return fahrenheit(c*9/5 + 32)
// }

// func calibrate(c celsiusF, offset celsius) celsiusF {
// 	return func() celsius {
// 		return c() + offset
// 	}
// }

// func main() {

// 	fmt.Println("==========================")
// 	fmt.Println("| °C         | °F        |")
// 	fmt.Println("==========================")

// 	var offset celsius = 5
// 	for i := 0; offset < 65; i++ {
// 		c := calibrate(getCelsius, offset)
// 		tempC := c()
// 		tempF := getFahrenheit(tempC)

// 		fmt.Printf("| %6.2f °C | %6.2f  °F |\n", tempC, tempF)
// 		offset += 5
// 	}
// }

// Мой вариант №2
type celsius float64
type fahrenheit float64
type kelvin float64

const (
	nameOperation = "%v\n"
	line          = "==========================="
	rowFormat     = "| %10.8v | %10.8v |\n"
	numberFormat  = "%.1f %v"
)

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9/5 + 32)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

type getRowFn func(row int, col1 string, col2 string) (string, string)

func drawTable(header, col1, col2 string, rows int, getRow getRowFn) {
	fmt.Printf(nameOperation, header)
	fmt.Println(line)
	fmt.Printf(rowFormat, col1, col2)
	fmt.Println(line)
	for i := 0; i < rows; i++ {
		c1, c2 := getRow(i, col1, col2)
		fmt.Printf(rowFormat, c1, c2)
	}
	fmt.Println(line)
}

func ctof(row int, c1 string, c2 string) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	col1 := fmt.Sprintf(numberFormat, c, c1)
	col2 := fmt.Sprintf(numberFormat, f, c2)
	return col1, col2
}

func ftoc(rows int, c1 string, c2 string) (string, string) {
	f := fahrenheit(rows*5 + 40)
	c := f.celsius()
	col1 := fmt.Sprintf(numberFormat, f, c1)
	col2 := fmt.Sprintf(numberFormat, c, c2)
	return col1, col2

}

func ktoc(row int, c1, c2 string) (string, string) {
	k := kelvin(row*5) + 235
	c := k.celsius()
	col1 := fmt.Sprintf(numberFormat, k, c1)
	col2 := fmt.Sprintf(numberFormat, c, c2)
	return col1, col2
}

func main() {
	drawTable("Celsius to Fahrenheit", "°C", "°F", 10, ctof)
	drawTable("Fahrenheit to Celsius", "°F", "°C", 5, ftoc)
	drawTable("Kelvin to Celsius", "K", "°C", 10, ktoc)
}

// Пиример
// type celsius float64

// func (c celsius) fahrenheit() fahrenheit {
// 	return fahrenheit((c * 9.0 / 5.0) + 32.0)
// }

// type fahrenheit float64

// func (f fahrenheit) celsius() celsius {
// 	return celsius((f - 32.0) * 5.0 / 9.0)
// }

// const (
// 	line         = "======================="
// 	rowFormat    = "| %8s | %8s |\n"
// 	numberFormat = "%.1f"
// )

// type getRowFn func(row int) (string, string)

// // drawTable создает таблицу с двумя столбцами.
// func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
// 	fmt.Println(line)
// 	fmt.Printf(rowFormat, hdr1, hdr2)
// 	fmt.Println(line)
// 	for row := 0; row < rows; row++ {
// 		cell1, cell2 := getRow(row)
// 		fmt.Printf(rowFormat, cell1, cell2)
// 	}
// 	fmt.Println(line)
// }

// func ctof(row int) (string, string) {
// 	c := celsius(row*5 - 40)
// 	f := c.fahrenheit()
// 	cell1 := fmt.Sprintf(numberFormat, c)
// 	cell2 := fmt.Sprintf(numberFormat, f)
// 	return cell1, cell2
// }

// func ftoc(row int) (string, string) {
// 	f := fahrenheit(row*5 - 40)
// 	c := f.celsius()
// 	cell1 := fmt.Sprintf(numberFormat, f)
// 	cell2 := fmt.Sprintf(numberFormat, c)
// 	return cell1, cell2
// }

// func main() {
// 	drawTable("°C", "°F", 29, ctof)
// 	fmt.Println()
// 	drawTable("°F", "°C", 29, ftoc)
// }
