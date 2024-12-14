package main

import (
	"fmt"
	"strings"
)

func main() {
	// task_1()
	task_2()
}

// //////// task_1 /////////
func sourceGopher(downstream chan string) {
	for _, v := range []string{"a", "b", "b", "c", "d", "d", "d", "e"} {
		downstream <- v
	}
	close(downstream)
}

func removeDuplicates(upstream, downstream chan string) {
	prev := ""
	for v := range upstream {
		if v != prev {
			downstream <- v
			prev = v
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func task_1() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go removeDuplicates(c0, c1)
	printGopher(c1)
}

/////////////////////

// ///////////// task_2 /////////
func task_2() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher_2(c0)
	go splitWords(c0, c1)
	printGopher_2(c1)
}

func sourceGopher_2(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {

		downstream <- v
	}
	close(downstream)
}

func splitWords(upstream, downstream chan string) {
	for v := range upstream {
		for _, word := range strings.Fields(v) {
			downstream <- word
		}
	}
	close(downstream)
}

func printGopher_2(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

//////////////////////////////////////
