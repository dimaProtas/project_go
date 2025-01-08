package main

import "fmt"

func main() {
	artur := &character{name: "Artur"}
	putnic := &character{name: "Putnic"}
	chery := &item{name: "шит"}
	mech := &item{name: "меч"}
	artur.pickup(mech)
	artur.pickup(chery)
	artur.givent(putnic)
	fmt.Print(artur)
	fmt.Print(putnic)
}

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if i == nil {
		fmt.Printf("%v ничего не поднял!\n", c.name)
		return
	}
	if c.leftHand != nil {
		fmt.Printf("%v не может взять %v, у него заняты руки он несет %v!\n", c.name, i.name, c.leftHand.name)
		return
	}
	fmt.Printf("%v поднимает %v\n", c.name, i.name)
	c.leftHand = i
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v несет ничего!\n", c.name)
	}
	return fmt.Sprintf("%v несет %v!\n", c.name, c.leftHand.name)
}

func (c character) givent(to *character) {
	if c.leftHand == nil {
		fmt.Printf("%v ничего не может дать %v\n", c.name, to.name)
		return
	}
	fmt.Printf("%v отдает %v - %v\n", c.name, c.leftHand.name, to.name)
	to.leftHand = c.leftHand
	c.leftHand = nil
}
