package main

import (
	"fmt"
	"sort"
)

// Вызывает ли nil сбои в Golang?
func task_1() {
	var nowhere *int
	fmt.Println(nowhere)
	// fmt.Println(*nowhere) // выводит ошибку так как указаьель никуда не указывает!
	if nowhere != nil { // избежание ошибки с разыменованием nil
		fmt.Println(*nowhere)
	}
}

// Защита методов в Golang
func task_2() {
	var nobody *persson
	fmt.Println(nobody)

	nobody.birthday()
}

type persson struct {
	age int
}

func (p *persson) birthday() {
	if p != nil {
		p.age++
	} else {
		fmt.Println("Приёмник равен nill!\nВо избежании ошибки прописан блок if!")
		return
	}
}

// Значения функций nil в Golang
func task_3() {
	// Когда переменная объявляется как тип функции, ее значение по умолчанию будет nil.
	var fn func(a, b int) int
	fmt.Println(fn)

	food := []string{"onion", "carrot", "celery"}

	sortStrings(food, nil) // сортировка по алфавиту
	fmt.Println(food)
	sortStrings(food, func(i, j int) bool { return len(food[i]) < len(food[j]) }) // сортировка по длине
	fmt.Println(food)
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i int, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

// Срезы nil в Golang
func task_4() {
	// Срез, объявленный без композитного литерала или встроенной функции make, будет со значением nil.
	// К счастью, ключевое слово range, а также встроенные функции len и append будут работать со срезами nil, как показано в следующем примере.
	var soup []string
	fmt.Println(soup == nil)

	for _, s := range soup {
		fmt.Println(s) // Ничего не выводит так как масив равен nill
	}

	fmt.Println(len(soup))

	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)

	masiv := mirepoix(nil)
	fmt.Println(masiv)
}

// Если вы пишите функцию, что принимает срез, убедитесь, что срез nil ведет себя так же, как и пустой срез.
func mirepoix(ingridients []string) []string {
	return append(ingridients, "onion", "carrot", "celery")
}

// Карты nil в Golang
func task_5() {
	var soup map[string]int
	// soup := map[string]int{
	// 	"cery":  1,
	// 	"apple": 2,
	// }
	fmt.Println(soup == nil)

	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}

	for ingredient, measurement := range soup {
		fmt.Println(ingredient, measurement)
	}
	// Если функция только читает из карты, можно передавать nil функции вместо создания пустой карты.
}

// Интерфейсы nil в Go
func task_6() {
	// Когда объявляется переменная типа интерфейса без присваивания, нулевым значением является nil.
	// Следующий листинг показывает, что тип интерфейса и значение являются nil, и переменная считается равной nil.
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	// Когда переменной с типом интерфейса присваивается значение, интерфейс внутренне указывает на тип и значение данной переменной.
	// Это приводит к довольно удивительному поведению значения nil, что не считается равным nil.
	// Оба, тип интерфейса и значение должны быть типа nil для того чтобы переменная была равна nil, что показано в следующем примере:
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)
	// Специальный символ %#v  нужен для того, чтобы увидеть тип и значение, а также отображения того,
	// что переменная содержит (*int)(nil), а не просто <nil>, как показано в Листинге 12.
	fmt.Printf("%#v\n", v)
	// Во избежание неприятных сюрпризов при сравнении интерфейсов с nil лучше явно использовать идентификатор nil вместо указания на переменную,
	// что содержит nil.
}

// Альтернатива nil в Golang
func task_7() {
	// Если значения нет, использование nil может показаться заманчивым. К примеру, указатель на целое число (*int) может представлять как ноль,
	// так и nil. Указатели нужны для указания, поэтому использование указателя просто для предоставления значение nil не является лучшим вариантом.

	// Вместо использования указателя можно объявить небольшую структуру с несколькими методами. Это требует большего количества кода,
	// но не запрашивает указатель или nil, как показано в следующем листинге.
	n := newNumber(45)
	fmt.Printf("%s - %s\n", n, n.String())

	e := number{}
	fmt.Printf("%s\n", e)
}

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		fmt.Printf("Valid равен %v!\n", n.valid)
	}
	return fmt.Sprintf("%d", n.value)
}

// Пример преподователя
func task_8() {
	arthur := &character{name: "Артур"}

	shrubbery := &item{name: "кустарник"}
	arthur.pickup(shrubbery) // Выводит: Артур поднимает кустарник

	knight := &character{name: "Рыцарь/ю"}
	arthur.give(knight) // Выводит: Артур дает Рыцарь/ю кустарник

	fmt.Println(arthur) // Выводит: Артур ничего не несет
	fmt.Println(knight) // Выводит: Рыцарь/ю несет кустарник
}

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v поднимает %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v ничего не может дать\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v с занятыми руками\n", to.name)
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v дает %v %v\n", c.name, to.name, to.leftHand.name)
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v ничего не несет", c.name)
	}
	return fmt.Sprintf("%v несет %v", c.name, c.leftHand.name)
}

func main() {
	// task_1() // Вызывает ли nil сбои в Golang?
	// task_2() // Защита методов в Golang
	// task_3() // Значения функций nil в Golang
	// task_4() // Срезы nil в Golang
	// task_5() // Карты nil в Golang
	// task_6() // Интерфейсы nil в Go
	// task_7() // Альтернатива nil в Golang
	task_8() // Пример преподователь DZ
}
