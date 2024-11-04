package main

import (
	"fmt"
)

// Создание своих собственных типов данных в Golang
type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Printf("В %v келинах - %2.4v °C градусов Цельсия\n", k, c)
}
