package main

import (
	"fmt"
	"math/rand"
	"time"
)

type zoo string

type Animals struct {
	zebra  zoo
	gepard zoo
	volk   zoo
}

type Stringer interface {
	String() string
}

func (z zoo) String() string {
	return fmt.Sprintf("Марсиансикй зверь - %v", string(z))
}

func (z zoo) move() string {
	switch rand.Intn(3) {
	case 0:
		return fmt.Sprintf("%v бежит\n", z)
	case 1:
		return fmt.Sprintf("%v сидит\n", z)
	default:
		return fmt.Sprintf("%v спит\n", z)
	}
}

func (z zoo) et() string {
	e := [...]string{"траву", "мясо", "молоко", "яблоки"}
	num := rand.Intn(len(e))
	return fmt.Sprintf("%v Кушает %v\n", z, e[num])
}

func main() {
	animals := Animals{
		zebra:  zoo("Gepard"),
		gepard: zoo("Zebra"),
		volk:   zoo("Volk"),
	}

	animalsSlise := [...]zoo{animals.zebra, animals.gepard, animals.volk}

	for i := 0; i <= 72; i++ {
		rundInt := rand.Intn(len(animalsSlise))
		dayTime := i % 24 // Определяем часы текущего дня
		if i <= 23 {
			fmt.Println("День 1")
		} else if i >= 48 {
			fmt.Println("День 3")
		} else {
			fmt.Println("День 2")
		}
		switch {
		case dayTime >= 6 && dayTime < 22:
			fmt.Printf("Час %d: Это день. \n", dayTime)
			if i%2 == 0 {
				fmt.Println(animalsSlise[rundInt].move())
			} else if i%2 == 1 {
				fmt.Println(animals.volk.et())
			}
		default:
			fmt.Printf("Час %d: Это ночь. \n", dayTime)
			fmt.Printf("Всеживотные спят!\n")
		}
		time.Sleep(time.Second)
	}

	// fmt.Println(animals.gepard.String())
	// // Движение зверей
	// fmt.Println(animals.zebra.runAnimals())
	// fmt.Println(animals.volk.sits())

	// // Кормление зверей
	// fmt.Println(animals.volk.et())
}
