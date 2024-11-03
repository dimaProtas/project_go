package main

import (
	"fmt"
)

// Объявление новых типов Golang
// func main() {
// 	type celsius float64
// 	var dreges celsius = 20

// 	var temperature celsius = dreges

// 	temperature += 10

// 	fmt.Println(temperature)
// }

// func main() {
// 	type celsius float64
// 	type fahrenheit float64

// 	var c celsius = 20
// 	var f fahrenheit = 20
// 	// Ошибка в операции: несовпадение типов celsius и fahrenheit
// 	// if c == f {
// 	// 	fmt.Println("true")
// 	// } else {
// 	// 	fmt.Println("false")
// 	// }
// 	// c += f

// }

// Создание своих собственных типов данных в Golang
// type celsius float64
// type kelvin float64

// func kelvinToCelsius(k kelvin) celsius {
// 	return celsius(k - 273.15)
// }

// func main() {
// 	var k kelvin = 294.0
// 	c := kelvinToCelsius(k)
// 	fmt.Printf("В %v келинах - %2.4v °C градусов Цельсия\n", k, c)
// }

// Напишите функцию celsiusToKelvin, что использует типы celsius и kelvin, определенные в Листинге 4. Используйте ее для конвертации 127° C,
// температуры освещенной солнцем поверхности луны, в градусы Кельвина.
// type celsius float64
// type kelvin float64

// func celsiusToKelvin(c celsius) kelvin {
// 	return kelvin(c + 273.15)
// }

// func main() {
// 	var c celsius = 127
// 	k := celsiusToKelvin(c)
// 	fmt.Printf("В %v°C - %v градусов Келвина!", c, k)
// }

// Добавление поведения типам через методы Go
// type kelvin float64
// type celsius float64
// type fahrenheit float64

// func (k kelvin) celsius() celsius {
// 	return celsius(k - 273.15)
// }

// func (f fahrenheit) celsius() celsius {
// 	return celsius((f - 32.0) * 5.0 / 9.0)
// }

// func main() {
// 	var k kelvin = 294.0
// 	c := k.celsius()
// 	fmt.Printf("В %v°К Келвинах - %2.4v°C градусов Цельсия\n", k, c)

// 	var f fahrenheit = 20.0
// 	c_f := f.celsius()
// 	fmt.Printf("В %v°F - %2.4v°C градусов Wельсия\n", f, c_f)
// }

// Напишите программу с типами celsius, fahrenheit и kelvin и методами для конвертации из одного типа температуры в другой.
type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((k-273.15)*9.0/5.0 + 32.0)
}
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}
func (f fahrenheit) kelvin() kelvin {
	return kelvin((f-32.0)*5.0/9.0 + 273.15)
}
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var c celsius = 27.0
	var f fahrenheit = 68.0
	var k kelvin = 294.0

	k_f := k.fahrenheit()
	c_f := c.fahrenheit()
	c_k := c.kelvin()
	f_k := f.kelvin()
	f_c := f.celsius()
	k_c := k.celsius()

	fmt.Printf("В %v°C - %2.4v°F\n", c, c_f)
	fmt.Printf("В %v°K - %2.4v°F\n", k, k_f)
	fmt.Printf("В %v°F - %2.4v°C\n", f, f_c)
	fmt.Printf("В %v°K - %2.4v°C\n", k, k_c)
	fmt.Printf("В %v°C - %2.4v°K\n", c, c_k)
	fmt.Printf("В %v°F - %2.4v°K\n", f, f_k)
}
