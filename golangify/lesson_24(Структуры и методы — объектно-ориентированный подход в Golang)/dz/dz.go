package main

import (
	"fmt"
	"math"
)

type location struct {
	modulus   string
	name      string
	lat, long float64
}

type locationRes struct {
	modulus, name, lat, long string
}

func conversionToDMS(value float64) string {
	// Отделяем градусы
	degrees := int(value)
	// Учитываем знак для градусов
	if degrees < 0 {
		value = -value
		degrees = -degrees
	}

	// Вычисляем минуты
	minutesFloat := (math.Abs(value) - math.Abs(float64(degrees))) * 60
	minutes := int(minutesFloat)

	// Вычисляем секунды
	seconds := (minutesFloat - float64(minutes)) * 60

	return fmt.Sprintf("%d° %d' %.2f''", degrees, minutes, seconds)
}

func decimalDMT(l location) locationRes {
	lat := conversionToDMS(l.lat)
	long := conversionToDMS(l.long)

	return locationRes{l.modulus, l.name, lat, long}
}

func DZ1() {
	spirit := location{"Spirit", "Columbia Memorial Station", -14.5684, 175.472636}
	opportunity := location{"Opportunity", "Challenger Memorial Station", -1.9462, 354.4734}
	curiosity := location{"Curiosity", "Brandbury Landing", -14.5684, 175.472636}
	insight := location{"Elysium Planitia", "", 4.5, 135.9}

	locations := []location{spirit, opportunity, curiosity, insight}

	fmt.Printf("%-30v %-30v %-30v %-30v\n", "Модуль", "Место посадки", "Широта", "Долгота")
	fmt.Println("===============================================================================================================")

	for i := 0; i < len(locations); i++ {
		res := decimalDMT(locations[i])
		fmt.Printf("%-30v %-30v %-30v %-30v\n", res.modulus, res.name, res.lat, res.long)
	}

}
