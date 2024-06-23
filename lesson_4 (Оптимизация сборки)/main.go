package main

func main() {
	println("Hellow world")
}

// Компиляция ```go build main.go```
// Запуск в linux ./main
// Оптимизированная сборка файла ```go build -o new_main -ldflags "-s -w" main.go```` (занимает меньше памяти)
