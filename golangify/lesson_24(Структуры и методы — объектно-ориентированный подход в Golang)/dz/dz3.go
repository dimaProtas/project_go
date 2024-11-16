package main

import (
	"fmt"
	"math"
)

// Найдите расстояние от вашего города до столицы вашей страны;
func DZ3() {
	earth := world{6371.8}
	mars := world{3959.8}
	// volkovisk := location2{"Volkovisk", "Belarus", 53.161277, 24.473064}
	// moscow := location2{"Moskow", "Russia", 55.7522, 37.6156}

	volkovisk := NewLocation(Coordinate{53, 10, 22.2, 'S'}, Coordinate{24, 28, 30.12, 'W'})
	moscow := NewLocation(Coordinate{55, 44, 24.00, 'S'}, Coordinate{37, 36, 36.00, 'W'})

	dist := earth.distance2(volkovisk, moscow)
	fmt.Printf("Растояние от Volkovisk до Moskow равно %.2f\n", dist)

	mountSharp := NewLocation(Coordinate{5, 4, 48.00, 'S'}, Coordinate{137, 51, 0, 'E'})
	OlimpusMons := NewLocation(Coordinate{18, 39, 0, 'N'}, Coordinate{226, 12, 0, 'E'})
	distMars := mars.distance2(mountSharp, OlimpusMons)
	fmt.Printf("Растояние от Mount Sharp до Olimpus Mons равно %.2f\n", distMars)

}

type location3 struct {
	lat, long float64
}

func NewLocation(Lat, Long Coordinate) location3 {
	return location3{Lat.decimal(), Long.decimal()}
}

type Coordinate struct {
	d, m, s float64
	h       rune
}

func (c Coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (w world) distance2(p1, p2 location3) float64 {
	// Добавить математические выражения, используя w.radius
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}
