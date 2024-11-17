package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("new_file.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Запись в фаил
	newFile, err2 := os.Create("new.csv")
	if err2 != nil {
		panic(err2)
	}
	defer newFile.Close()

	// Карта для отслеживания уникальных строк
	unqiq := make(map[string]bool)

	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		str := strings.Split(line, ",")
		if str[1] == "PC" {

			record := fmt.Sprintf("%s, %s\n", str[0], str[1])

			if !unqiq[record] { // проверям есть ли вкарте эта строка!
				unqiq[record] = true
				_, err3 := fmt.Fprintf(newFile, "%s, %s\n", str[0], str[1])
				// fmt.Printf("%s ------ %s\n", str[0], str[1])
				if err3 != nil {
					panic(err3)
				}
			}

		}
	}
	// println(lineCount)
}
