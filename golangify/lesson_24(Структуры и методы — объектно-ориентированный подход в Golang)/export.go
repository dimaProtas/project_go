package main

import "math"

// Импортируеться в main.go для task_1 и task_2
type Location struct {
	Lat, Long float64
}

type Coordinate struct {
	d, m, s float64
	h       rune
}

func (c Coordinate) Decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func NewLocation(Lat, Long Coordinate) Location {
	return Location{Lat.Decimal(), Long.Decimal()}
}

// для task_3
type World struct {
	radius float64
}

func (w World) Distance(p1, p2 Location) float64 {
	// Добавить математические выражения, используя w.radius
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	clong := math.Cos(rad(p1.Long - p2.Long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}

// rad конвертирует градусы в радианы.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}
