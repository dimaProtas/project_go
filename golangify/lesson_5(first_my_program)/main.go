package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Запись в фаил
func main() {
	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	fmt.Printf("%-18v %-4v %8v %3v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Printf("%s\n", strings.Repeat("=", 40))
	_, err = fmt.Fprintf(file, "%-18v, %-4v, %8v, %3v\n", "Spaceline", "Days", "Trip type", "Price")

	for i := 0; i < 10; i++ {
		var spaceline string
		speed := rand.Intn(15) + 16
		distance := 62100000
		secondsPerDay := 85400
		duration := distance / speed / secondsPerDay

		switch num := rand.Intn(3); num {
		case 0:
			spaceline = "SpaceX"
		case 1:
			spaceline = "Space Adventures"
		case 2:
			spaceline = "Virgin Galactic"

		}
		var type_tripe string
		price := 20.0 + speed

		if rand.Intn(2) == 0 {
			type_tripe = "Round-trip"
			price = price * 2
		} else {
			type_tripe = "One-way"
		}

		_, err = fmt.Fprintf(file, "%-18v, %-4v, %-10v, $%4v\n", spaceline, duration, type_tripe, price)
		fmt.Printf("%-18v %-4v %-10v $%4v\n", spaceline, duration, type_tripe, price)
	}
	defer file.Close()

}
