package main

import (
	"fmt"
	"strings"
	"time"
)

// Амперсанд (&) и звездочка астериск (*) в Golang
func task_1() {
	// Это место в памяти, где компьютер хранит 42. К счастью, вы можете использовать название переменной answer
	// для получения переменной вместо того адреса памяти, что использует компьютер.
	answer := 42
	fmt.Println(&answer)

	// Оператор адреса (&) предоставляет адрес памяти значения. Обратная операция называется разыменованием, что выдает значение, к которому обращается
	// адрес памяти. В следующем листинге происходит разыменование переменная address разыменуется через префикс в виде астерикса (*).
	address := &answer
	fmt.Printf("Разименовывание переменной: %v\n", *address)
}

// Типы указателей в Golang
func task_2() {
	// Переменная address в Листинге 2 является указателем типа *int, об этом сообщает специальный символ %Т в следующем примере.
	answer := 42
	adddress := &answer
	fmt.Printf("Тип указателя: %T\n", adddress)

	// Типы указателя могут появиться везде, где типы используются, включая объявления переменных, параметры функции, возвращаемые типы,
	// типы полей структуры и так далее. В следующем примере звездочка (*) в объявлении home поясняет, что это тип указателя.
	canda := "Canada"
	// Звездочка перед типом обозначает тип указателя, а звездочка перед названием переменной нужна для указания на значение, к которому отсылается указатель.
	var home *string

	fmt.Printf("Тип указателя: %T\n", home)

	home = &canda
	fmt.Printf("Разименовываем указатель: %v\n", *home)
}

// Указатели для указания Golang
func task_3() {
	// Изменить значение bolden можно в одном месте, потому что переменная administrator указывает на bolden вместо хранения копии:
	var administrator *string

	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	// Можно разыменовать administrator для непрямого изменения значения bolden:
	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(*administrator)
	fmt.Println(bolden)

	// Результатом присваивания major к administrator является новый указатель, что также указывает на строку bolden.
	mojor := administrator
	*mojor = "NewAdministrator"
	fmt.Println(bolden)
	// Указатели major и administrator оба содержат один и тот же адрес памяти, следовательно, они равны:
	fmt.Printf("Проверка равен ли указатель и переменная: %v\n", administrator == mojor)

	// 20 января 2017 года на смену Чарльзу Болдену пришел Роберт М. Лайтфут. После данного изменения administrator и major перестали указывать на одинаковый адрес памяти
	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(*administrator)
	fmt.Printf("Проверка равен ли указатель и переменная: %v\n", administrator == mojor)

	// Присваивание разыменованного значения major к другой переменной создает копию строки.
	// После создания клона прямые и непрямые изменения с bolden не будут иметь эффект над значением charles и наоборот:
	charles := *mojor
	*mojor = "Charles Bolden"
	fmt.Println(charles)
	fmt.Println(bolden)

	// Если две переменные содержат одинаковую строку, они считаются равными, как в случае с charles и bolden в следующем коде.
	// Даже несмотря на то, что их адреса памяти отличаются:
	charles = "Charles Bolden"
	fmt.Printf("Проверка равны ли строки: %v\n", charles == bolden)
	fmt.Printf("Проверка павны ли указатели: %v\n", &charles == &bolden)

}

// Указатели и структуры Go
func task_4() {
	// В отличие от строк и чисел, перед композитными литералами ставится префикс в виде оператора адреса.
	// В следующем примере переменная timmy содержит адрес памяти, указывающий на структуру person.
	type person struct {
		name, superpower string
		age              int
	}

	timmy := &person{
		name: "Timmy",
		age:  25,
	}

	fmt.Printf("%+v\n", timmy)
	fmt.Println(timmy)

	// Кроме того, нет необходимости разыменования структур для получения доступа к их полю. Следующий листинг предпочтителен для написания (*timmy).superpower.
	timmy.superpower = "Flying"
	fmt.Println(timmy)
	fmt.Printf("%+v\n", timmy)
}

// Указатели и массивы в Go
func task_5() {
	// Как и в случае со структурами, композитные литералы для массивов могут дополнены префиксом в виде оператора адреса (&) для создания нового
	// указателя на массив. Массивы также предоставляют автоматическое разыменование, как показано в следующем примере.
	superpower := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Printf("Автоматическое разименовывание масива: %v\n", superpower[0])
	fmt.Println(superpower[1:2])
	fmt.Println(superpower[2:])

	// Композитным литералам для срезов и карт также можно добавить префиксы с оператором адреса (&), однако тогда не будет автоматического разыменования.
	tetsMap := map[string]int{"one": 1, "two": 2}
	fmt.Println(tetsMap["one"])
	// В Go карты уже работают как ссылочный тип, поэтому передавать карту через указатель (&map) нет необходимости.
}

// Указатели в качестве параметров Golang
type person struct {
	name, superpower string
	age              int
}

func task_6() {
	// В Листинге 9 функция birthday объявляется с одним параметром типа *person. Это позволяет телу функции разыменовывать указатель и редактировать
	// значение, на которое он указывает. Как и в Листинге 7, здесь не обязательно в открытую разыменовывать переменную p для получения доступа к
	// полю age. Синтаксис следующего листинга предпочтительнее (*p).age++.

	dima := &person{
		name: "Dima",
		age:  30,
	}

	fmt.Printf("Выводим структуру dima: %+v\n", dima)
	birthday(dima)
	fmt.Printf("dima после фнкции birthday: %+v\n", dima)

	rebeca := person{
		name:       "Rebecca",
		superpower: "speed",
		age:        25,
	}

	fmt.Printf("Выводим структуру rebeca: %+v\n", rebeca)
	birthday(&rebeca)
	fmt.Printf("rebeca после фнкции birthday: %+v\n", rebeca)

}
func birthday(p *person) {
	p.age += 2
}

// Приемники указателя в Golang
// Вызывается с указателем или нет, но метод birthday, объявленный в Листинге 11, должен уточнить приемник указателя — в противном случае age не увеличится.
func (p *person) birthday() { // если не поставить * то значение приемника (приемник p) не увеличиться
	p.age += 2
}

func task_7() {

	ira := person{
		name:       "Ira",
		superpower: "Love",
		age:        25,
	}

	fmt.Printf("Выводим структуру ira: %+v\n", ira)
	ira.birthday()
	fmt.Printf("ira после метода birthday: %+v\n", ira)

	// Структуры регулярно передаются с указателями. Для метода birthday будет иметь смысл изменить атрибуты существующей person вместо создания новой
	// персоны. Тем не менее, не каждую структуру стоит изменять. Стандартная библиотека предоставляет отличный пример в виде пакета time.
	// Методы типа time.Time никогда не используют приемник указателя, предпочитая возвращать вместо этого новое время, как показано в следующем примере.
	// В конечном итоге, завтра — это новое сегодня.
	const layout = "Mon, Jan 2, 2006"
	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)
	fmt.Printf("%v\n", day.Format(layout))
	fmt.Printf("%v\n", tomorrow.Format(layout))
}

// Внутренние указатели Golang
func task_8() {
	dima := stats{
		level:     1,
		endurance: 42,
		health:    5,
	}

	fmt.Println(dima)

	// У типа character нет указателей в определении структуры, однако в случае необходимости вы можете взять адрес памяти любого поля.
	// Код &player.stats предоставляет указатель на внутреннюю часть структуры.
	player := character{
		name: "Dima",
	}

	lavelUp(&player.stats)
	fmt.Printf("%+v\n", player.stats)
}

type character struct {
	name  string
	stats stats
}

type stats struct {
	level             int
	endurance, health int
}

func lavelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance

}

// Изменение, или мутации массивов в Golang
func task_9() {
	// Хотя предпочтительнее использовать срезы, а не массивы, в случаях, когда нет нужды менять их длину, можно задействовать и массивы. В качестве примера можно рассмотреть шахматную доску.
	// В следующем примере показано, как указатели позволяют функциям изменять элементы массива.
	var board [8][8]rune
	reset(&board)

	fmt.Printf("%c\n", board)

	// Работа со срезами
	var list [][]int
	list = append(list, make([]int, 4))
	newFunc(list)

	fmt.Println(list)

}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func newFunc(l [][]int) {
	l[0][1] = 34
}

// Карты в роли указателей в Go
func task_10() {
	// В уроке про карты мы говорили о том, что во время присваивания или передачи в качестве аргументов карты не копируются. \
	// Карты являются скрытыми указателями, поэтому указание на карту является лишним. Не делайте этого:
	// func demolish(planets *map[string]string) // Лишний указатель
}

// Срезы в Go — указатели на массив
func task_11() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}

	replaysify(&planets)

	fmt.Println(planets)
}

// Явный указатель на срез полезен только в том случае, когда нужно модифицировать сам срез: длину, вместимость или начальный набор.
// В следующем примере функция reclassify редактирует длину среза planets. Вызов функции (main) не увидит изменения,
// если reclassify не использует указатель.
func replaysify(planets *[]string) {
	*planets = (*planets)[0:8]
}

// Указатели и интерфейсы в Golang
func task_12() {
	shout(martian{})
	shout(&martian{})

	pew := laser(2)
	shout(&pew)
}

// В следующем примере показано, что martian и указатель на martian удовлетворяют интерфейсу talker.
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

// Все иначе, когда методы используют приемники указателя, как показано далее:
type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main() {
	// task_1() // Амперсанд (&) и звездочка астериск (*) в Golang
	// task_2() // Типы указателей в Golang
	// task_3() // Указатели для указания Golang
	// task_4() // Указатели и структуры Go
	// task_5() // Указатели и массивы в Go
	// task_6() // Указатели в качестве параметров Golang
	// task_7() // Приемники указателя в Golang
	// task_8() // Внутренние указатели Golang
	// task_9() // Изменение, или мутации массивов в Golang/
	// task_10() // Карты в роли указателей в Go
	// task_11() // Срезы в Go — указатели на массив
	task_12() // Указатели и интерфейсы в Golang
}
