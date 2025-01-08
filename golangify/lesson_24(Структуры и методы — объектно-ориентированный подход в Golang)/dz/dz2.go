package main

import (
	"fmt"
	"math"
)

type location2 struct {
	modules, name string
	lat, long     float64
}

type world struct {
	radius float64
}

// rad конвертирует градусы в радианы.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location2) float64 {
	// Добавить математические выражения, используя w.radius
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}

func task_1() {
	mars := world{3389.7}

	spirit := location2{"Spirit", "Columbia Memorial Station", -14.5684, 175.472636}
	opportunity := location2{"Opportunity", "Challenger Memorial Station", -1.9462, 354.4734}
	curiosity := location2{"Curiosity", "Brandbury Landing", -14.5684, 175.472636}
	insight := location2{"Elysium Planitia", "", 4.5, 135.9}

	locations := []location2{spirit, opportunity, curiosity, insight}

	for i := 0; i < len(locations); i++ {
		for j := 0; j < len(locations); j++ {
			if locations[i].name != locations[j].name {
				dist := mars.distance(locations[i], locations[j])
				fmt.Printf("Растояние от %v до %v равно %.2f\n", locations[i].modules, locations[j].modules, dist)
			}
		}
	}
}

// Найдите расстояние от Лондона, Англия (London, England, 51°30’N 0°08’W) до Парижа, Франция (Paris, France, 48°51’N 2°21’E);
func task_2() {
	earth := world{6371.8}
	london := location2{"London", "England", 51.509865, -0.118092}
	paris := location2{"Paris", "France", 48.856614, 2.35222}

	dist := earth.distance(london, paris)
	fmt.Printf("Растояние от %v до %v равно %.2f\n", london.modules, paris.modules, dist)
}

func DZ2() {
	// 1. Какие места посадки находятся ближе всего друг к другу?
	// 2. Какие места дальше всего друг от друга?
	// task_1()
	// Найдите расстояние от Лондона, Англия (London, England, 51°30’N 0°08’W) до Парижа, Франция (Paris, France, 48°51’N 2°21’E);
	task_2()
}
