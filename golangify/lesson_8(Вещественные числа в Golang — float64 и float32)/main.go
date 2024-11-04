package main

import (
	"fmt"
	"math/rand"
)

func task_1() {
	money := 0.0
	n := 0
	for n < 11 {
		money += 0.10
		n += 1
	}
	fmt.Printf("%2.4f\n", money)
}

func task_2() {
	var money float32
	for i := 0; i < 11; i++ {
		money += 0.10
	}
	fmt.Printf("%2.20f\n", money)
}

func task_3() {
	pingBang := 0.0
	for pingBang < 20.00 {
		n := rand.Intn(3) + 1
		switch n {
		case 1:
			pingBang += 0.10
			fmt.Printf("%05.2f\n", pingBang)
		case 2:
			pingBang += 0.05
			fmt.Printf("%05.2f\n", pingBang)
		case 3:
			pingBang += 0.25
			fmt.Printf("%05.2f\n", pingBang)
		}
	}
	fmt.Println("=============== ==========")
	fmt.Printf("%15v %10.2f\n", "Вы накопили:", pingBang)
	fmt.Println("=============== ==========")
	// fmt.Println(n)
}

func main() {
	// task_1()
	// task_2()
	task_3()
}
