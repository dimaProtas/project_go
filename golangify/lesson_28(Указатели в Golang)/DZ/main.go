package main

import "fmt"

type cordinate struct {
	y int
	x int
}

func (c *cordinate) vverh() string {
	c.y++
	return fmt.Sprintf("x: %v, y: %v\n", c.x, c.y)
}

func (c *cordinate) vniz() string {
	c.y--
	return fmt.Sprintf("x: %v, y: %v\n", c.x, c.y)
}

func (c cordinate) vpravo() string {
	c.x++
	return fmt.Sprintf("x: %v, y: %v\n", c.x, c.y)
}

func (c cordinate) vlevo() string {
	c.x--
	return fmt.Sprintf("x: %v, y: %v\n", c.x, c.y)
}

func main() {
	tur := cordinate{y: 2, x: 2}
	fmt.Println(tur.vpravo())
	fmt.Println(tur.vlevo())
	fmt.Println(tur.vlevo())
	fmt.Println(tur.vverh())
	fmt.Println(tur.vverh())
	fmt.Println(tur.vniz())
}
