package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Скучно, когда одинаковая строка кода повторяется из раза в раз. Напишите элемент конвейера (горутину),
	// что запоминает предыдущее значение и только отправляет значение на следующий этап конвейера, если оно отличается от того, что пришло ранее.
	// Чтобы все упростить, можете предположить, что первая строка никогда не бывает простой.
	// task_1()

	// Иногда проще оперировать словами, а не предложениями. Напишите конвейер, что принимает строки и разделяет их на слова
	// (можете использовать функцию Fields из пакета strings), а также отправляет все слова, одно за другим, на следующую стадию конвейера.
	task_2()
}

// //////// task_1 //////////
func source(stream chan string) {
	for _, i := range []string{"ap", "one", "one", "two", "two", "three", "four"} {
		stream <- i
	}
	close(stream)
}

func filter(sour, stream chan string) {
	var str string
	for {
		item, ok := <-sour
		if !ok {
			close(stream)
			return
		}
		if item != str {
			str = item
			stream <- item
		}
	}
}

func print(stream chan string) {
	for i := range stream {
		fmt.Println(i)
	}
}

func task_1() {
	c0 := make(chan string)
	c1 := make(chan string)
	go source(c0)
	go filter(c0, c1)
	print(c1)
}

//////////////////////////////////

// /////////////////////// task_2 ///////////////////
func task_2() {
	v0 := make(chan string)
	v1 := make(chan string)
	go sourseString(v0)
	go filterStringBig(v0, v1)
	printStringBig(v1)

}

func sourseString(upstream chan string) {
	strBig := []string{"Иногда проще оперировать словами", "а не предложениями", " Напишите конвейер", "что принимает строки и разделяет их на слова", "(можете использовать функцию Fields из пакета strings)", "а также отправляет все слова, одно за другим", "на следующую стадию конвейера."}
	for _, i := range strBig {
		upstream <- i
	}
	close(upstream) // закрываем канал
}

func filterStringBig(downstream, upstream chan string) {
	re := regexp.MustCompile(`[(),.]`) // Указываем символы для удаления
	for i := range downstream {
		fields := strings.Fields(i)
		for _, f := range fields {
			upstream <- re.ReplaceAllString(f, "")
		}
	}
	close(upstream)
}

func printStringBig(downstream chan string) {
	for i := range downstream {
		fmt.Println(i)
	}
}

////////////////////////////////////////////////////
