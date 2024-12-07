package main

import (
	"fmt"
	"math"
)

// Напишите программу со структурой gps для Системы Глобального Позиционирования (GPS). Данная структура должна состоять из текущей местности,
// места назначения и мира.

// Реализация метода description для типа location возвращает строку, содержащую название, широту и долготу.
// Тип world должен имплементировать метод для расстояния, используемый в уроке о структурах и методах.

// Прикрепите два метода к типу gps. Для начала прикрепите метод distance, что находит расстояние между текущей местностью и местом назначения.
// Затем реализуйте метод message, что возвращает строку с оставшимися километрами до пункта назначения.

// Финальным шагом станет создание структуры rover, что внедряет gps и написание функции main для тестирования всего созданного.
// Инициализируйте GPS для Марса с текущей локацией Bradbury Landing (-4.5895, 137.4417) и пунктом назначения Elysium Planitia (4.5, 135.9).
// Затем создайте элемент curiosity для марсохода и выведите его message (что встраивается в gps).

type gps struct {
	location
	world
}

type world struct {
	radius float64
}

type location struct {
	location_current
	location_end
}

type location_current struct {
	name      string
	lat, long float64
}

type location_end struct {
	name      string
	lat, long float64
}

type rover struct {
	gps
	nameRover string
}

func (l location) description() string {
	return fmt.Sprintf("Точка старта:  %v, Широта %.2f, Долгота %.2f\nТочка оканчания маршрута:  %v, Широта %.2f, Долгота %.2f/n", l.location_current.name, l.location_current.lat, l.location_current.long, l.location_end.name, l.location_end.lat, l.location_current.long)
}

func (g gps) description() string {
	return g.location.description()
}

func (w world) distance(l location) float64 {
	// Добавить математические выражения, используя w.radius
	s1, c1 := math.Sincos(rad(l.location_current.lat))
	s2, c2 := math.Sincos(rad(l.location_end.lat))
	clong := math.Cos(rad(l.location_current.long - l.location_end.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}

func (g gps) distance() float64 {
	return g.world.distance(g.location)
}

// rad конвертирует градусы в радианы.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (g gps) message(k int) string {
	dist := g.distance()
	res := dist - float64(k)
	return fmt.Sprintf("Прошли %v км, Осталось %.2f км до %v.\n", k, res, g.location.location_end.name)
}

func task_1() {
	gps1 := gps{
		location{
			location_current{
				"Москва",
				55.751244,
				37.618423,
			},
			location_end{
				"Санкт-Петербург",
				59.939095,
				30.315868,
			},
		},
		world{
			6371.0,
		},
	}

	// fmt.Println(gps1.description())
	// fmt.Println(gps1.world.distance(gps1.location))
	fmt.Println(gps1.description())
	fmt.Printf("Растояние от %v до %v равно %v\n", gps1.location.location_current.name, gps1.location.location_end.name, gps1.distance())
	fmt.Println(gps1.message(300))
}

func main() {
	mars := gps{
		location{
			location_current{"Bradbury Landing", -4.5895, 137.4417},
			location_end{"Elysium Planitia", 4.5, 135.9},
		},
		world{3389.5},
	}

	curiosity := rover{
		gps:       mars,
		nameRover: "Curiosity",
	}

	fmt.Println(curiosity.description())
	fmt.Println(curiosity.message(300))
	// task_1()
}
