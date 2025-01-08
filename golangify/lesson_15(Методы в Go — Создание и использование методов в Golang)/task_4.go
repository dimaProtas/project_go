package main

import (
	"fmt"
)

// Добавление поведения типам через методы Go
type kelvin float64
type celsius float64
type fahrenheit float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var k kelvin = 294.0
	c := k.celsius()
	fmt.Printf("В %v°К Келвинах - %2.4v°C градусов Цельсия\n", k, c)

	var f fahrenheit = 20.0
	c_f := f.celsius()
	fmt.Printf("В %v°F - %2.4v°C градусов Wельсия\n", f, c_f)
}
