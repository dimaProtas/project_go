package main

import (
	"fmt"
	"strings"
)

// Расшифровать строку шифром Вижнера, используя ключевое слово GOLANG
func main_decript(cipherText string) {
	// buk := 'a'
	// fmt.Printf("%v, %[1]c\n", buk)
	// cipherText := "CSOITEUIWUIZNSROCNKFD"
	// cipherText := "ABC"
	keyword := "GOLANG"
	keyindex := 0
	message := ""

	for i := 0; i < len(cipherText); i++ {
		symbol := cipherText[i] - 'A'
		k := keyword[keyindex] - 'A'
		// fmt.Printf("%v,-%v,%[2]c\n", symbol, k)

		symbol = (symbol-k+26)%26 + 'A'
		message += string(symbol)

		keyindex++
		keyindex %= len(keyword)
	}
	fmt.Println(message)
}

// Зашифровка строки шифром Вижнера, с использованием ключевого слова GOLANG
func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"
	index := 0
	message := ""

	plainText = strings.ToUpper(strings.Replace(plainText, " ", "", -1))

	for i := 0; i < len(plainText); i++ {
		c := plainText[i]
		if c >= 'A' && c <= 'Z' {
			c -= 'A'
			k := keyword[index] - 'A'

			// println(c, k)
			c = (c+k)%26 + 'A'
			// fmt.Println(c)

			index++
			index %= len(keyword)
		}
		message += string(c)
	}
	fmt.Println(message)
	main_decript(message)
}
