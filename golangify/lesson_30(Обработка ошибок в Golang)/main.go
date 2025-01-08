package main

// DZ - преподователя!
import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Исправление ошибок в Golang
func task_1() {
	// Чтение содержимого директории
	files, err := ioutil.ReadDir("../")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Printf("dir: %v - %v\n", file.IsDir(), file.Name())
	}

	// Создание директории
	// err := os.Mkdir("testDir", 0750)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

}

// Элегантная обработка ошибок в Golang
func task_2() {
	err := proverbs("errors.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close()
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	f.Close()
	return err

}

// Применяем defer — отложенные действия в Golang
func task_3() {
	err := proverbs_defer("new.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Вы можете отложить любую функцию или метод, и как множество возвращаемых значений, отсрочка нужна не для уточнения обработки ошибки.
// Она улучшает обработку ошибок, избавляясь от необходимости постоянно помнить об очистке.
// Благодаря defer, код для обработки ошибок может сфокусироваться только на своей задачи и больше ни о чем.

func proverbs_defer(name string) error {
	f, err := os.Create(name) // Создание файла
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	return err
}

// Креативная обработка ошибок в Golang
func task_4() {
	proverbsElegantErr("elegant.csv")
}

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // Пропускает запись, если раньше была ошибка
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbsElegantErr(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := &safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is better than a little dependency.")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don't communicate by sharing memory, share memory by communicating.")
	sw.writeln("Channels orchestrate; mutexes serialize.")

	return sw.err
}

// Новые ошибки в программе на Golang
func task_5() {
	var g Grid
	err := g.Set(9, 0, 5)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(g)
}

const rows, columns = 9, 9

type Grid [rows][columns]int8

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit") // Создание новой ошибки!
)

func (g *Grid) Set(row, column int, digit int8) error {
	if !isBounds(row, column) {
		return ErrBounds // Создание новой ошибки!
	}
	if !validDigit(digit) {
		return ErrDigit // Создание новой ошибки!
	}
	g[row][column] = digit
	return nil
}

func isBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

// Причины каждой ошибки в Go
func test_6() {
	var g Grid
	err := g.Set(0, 0, 10)
	if err != nil {
		// Если метод Set возвращает ошибку, вызывающая сторона может различить возможные ошибки и обрабатывать определенные ошибки по-разному,
		// как показано в следующем листинге. Вы можете сравнить ошибку, возвращаемую с переменными ошибки, используя == или оператор switch.
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Println("Les erreurs de paramètres hors limites.")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
	fmt.Println(g)

}

func validDigit(d int8) bool {
	if d < 1 || d > 9 {
		return false
	}
	return true
}

// Настраиваемые типы ошибок в Golang
func task_7() {
	// Каким бы полезным не был errors.New, иногда нужно, чтобы ошибки описывались не просто сообщением. Go достаточно свободен в этом плане.

	// Тип error является встроенным интерфейсом, как показано в следующем примере. Любой тип, что имплементирует метод Error() для возвращения строки,
	// неявно удовлетворяет интерфейс. В качестве интерфейса возможно создать новые типы ошибок.
	type error interface {
		Error() string
	}
}

// Множество ошибок в Golang
func task_8() {
	// Вместо возвращения одной ошибки за раз, метод Set может сделать несколько проверок и вернуть все ошибки сразу.
	// Тип SudokuError в Листинге 15 является срезом error.
	// Он удовлетворяет интерфейсу error с методом, что соединяет ошибки вместе в одну строку.
	var g Grid
	err := g.SetNew(0, 0, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(g)

}

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, e := range se {
		s = append(s, e.Error())
	}
	return strings.Join(s, ",")
}

func (g *Grid) SetNew(row, column int, digit int8) error {
	var errs SudokuError
	if !isBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit
	return nil
}

// Утверждение типа в Go
func task_9() {
	var g Grid
	err := g.SetNew(10, 0, 15)
	if err != nil {
		if errs, ok := err.(SudokuError); ok { // пытается конвертировать значение err из типа интерфейса error в конкретный тип SudokuError.
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}
}

// Принцип работы panic в Golang
func task_10() {
	// При обработке ошибок некоторые языки программирования сильно полагаются на исключения.
	// В Go нет исключений, но в нем есть похожий механизм, который называется panic.
	// При задействовании panic в программе происходит сбой. Это похоже на случаи необработанных исключений в других языках.
}

// Есть ли исключения в Golang?
func task_11() {
	// Исключения в других языках значительно отличаются от значений ошибок в Go. Это заметно как в поведении, так и в имплементации.

	// Если функция выдает исключение, и никто не собирается его перехватить, это исключение доходит до вызывающей функции,
	// затем до вызывающей ту функцию и так далее, пока достигает вершины стека вызовов (например, функция main).

	// Исключения — это стиль обработки ошибок, который можно считать включенным. Часто они не занимают код,
	// тогда как выбор обработки исключений может привлекать изрядное количество специализированного кода. Это связано с тем,
	// что вместо использования существующих возможностей языка исключения обычно имеют специальные ключевые слова,
	// такие как try, catch, throw, finally, raise, rescue, except и так далее.

	// Значения ошибок в Go предоставляют простую, гибкую альтернативу исключениям, которые могут помочь вам создать надежное программное обеспечение.
	// Игнорирование значений ошибок в Go — это сознательное решение, которое становится очевидным каждому, кто читает полученный код.

	// В чем преимущество значений ошибок Go по сравнению с исключениями?
	// Go подталкивает разработчиков к тому, чтобы они поняли причину ошибок, что приводит к более понятным программам,
	// в то время как исключения обычно игнорируются по умолчанию. Значения ошибок не требуют специальных ключевых слов,
	// делая их более простыми и в то же время гибкими.
}

// Как использовать panic в Golang
func task_12() {
	// panic("Я забыл свое полотенце")
	// На заметку: Хотя значения ошибок обычно предпочтительнее panic, panic часто лучше, чем os.Exit в том,
	// что panic запустит любую отсроченную функцию, а os.Exit этого делать не станет.
	// panic(test())

	// В некоторых ситуациях Go предпочтет panic вместо значений ошибок, это может быть деление на ноль:
	var zero int
	_ = 42 / zero // Runtime error: integer divide by zero - целое число делится на ноль

}

func test() string {
	return fmt.Sprintln("Тест запуска функции в panic!")
}

// Тонкости работы с panic в Golang
func test_13() {
	// Чтобы panic не привел к сбою программы, Go предоставляет функцию recover, что показано в Листинге 18.

	// Отсроченные функции выполняются перед возвращением функции, даже в том случае, если задействуется panic.
	// Если отсроченная функция вызывает recover, panic остановится, и программа продолжит выполняться.
	// В таком случае цель у recover похожа на catch, except и rescue в других языках.

	defer func() {
		if e := recover(); e != nil { // Приходит в себя после panic
			fmt.Println(e) // Выводит: Я забыл свое полотенце
		}
	}()

	panic("Я забыл свое полотенце") // Приводит к panic

	// Где может использоваться встроенная функция recover?
	// Только отсроченные функции могут использовать recover.
}

func main() {
	// task_1() // Исправление ошибок в Golang
	// task_2() // Элегантная обработка ошибок в Golang
	// task_3() // Применяем defer — отложенные действия в Golang
	// task_4() // Креативная обработка ошибок в Golang
	// task_5() // Новые ошибки в программе на Golang
	// test_6() // Причины каждой ошибки в Go
	// task_7() // Настраиваемые типы ошибок в Golang
	// task_8() // Множество ошибок в Golang
	// task_9() // Утверждение типа в Go
	// task_10() // Принцип работы panic в Golang
	// task_11() // Есть ли исключения в Golang?
	// task_12() // Как использовать panic в Golang
	test_13() // Тонкости работы с panic в Golang
}
