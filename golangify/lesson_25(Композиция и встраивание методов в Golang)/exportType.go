package main

import "math"

// Композиция структур в Golang
type Celsius float64

type Temperature struct {
	high, low Celsius
}

type Location struct {
	lat, long float64
	World
}

type Report struct {
	sol         int
	temperature Temperature
	location    Location
}

// ========================== //
// методы из task_1
// Создавая отчет о погоде из более мелких типов, можно организовать код путем прикрепления методов каждого типа.
// Например, чтобы рассчитать среднюю температуру, можно написать метод, подобный данному ниже:
func (t Temperature) Average() Celsius {
	return (t.high + t.low) / 2
}

func (r Report) AverageReportsTemperature() Celsius {
	return r.temperature.Average()
}

// ========================== //
// методы из task_2
// К счастью, Go может встроить метод за вас с помощью внедрения struct.
// Для внедрения типа в структуру тип уточняется без указания названия поля, как показано в следующем примере:
type Report2 struct {
	sol
	Temperature // Тип temperature встроен в report
	Location
}

// К любым методам, объявленным в типе sol, можно получить доступ через поле sol или через тип report:
type sol int

func (s sol) Days(s2 sol) int {
	days := int(s - s2)
	if days < 0 {
		days = -days
	}

	return days
}

// ========================== //
// методы из task_3
type World struct {
	radius float64
}

func (l Location) Days(l2 Location) float64 {
	// Добавить математические выражения, используя w.radius
	s1, c1 := math.Sincos(rad(l.lat))
	s2, c2 := math.Sincos(rad(l2.lat))
	clong := math.Cos(rad(l.long - l2.long))
	return l.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}

// rad конвертирует градусы в радианы.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (r Report2) Days(s2 sol) int {
	return r.sol.Days(s2)
}

// ========================== /

// ========================== //
