package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Объявление структуры в Golang
func task_1() {
	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Printf("долгота = %.2f, широта = %.2f\n", curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}

// Повторное использование структур с типами в Go
func task_2() {
	type location struct {
		lat  float64
		long float64
	}

	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734

	fmt.Println(spirit, opportunity)

	// Как можно адаптировать код из Листинга 1 для использования типа location для марсохода Curiosity в кратере Брэдбери?
	var curiosity location
	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Printf("Марсаход Curiosity находиться в Долготе = %.2f, Широте = %.2f\n", curiosity.lat, curiosity.long)
}

// Инициализация структур через композитные литералы Golang
func task_3() {
	type location struct {
		lat, long float64
	}

	opportunity := location{lat: -1.9462, long: 354.4734}
	insight := location{lat: 4.5, long: 135.9}

	fmt.Println(opportunity, insight)

	// Композитный литерал в Листинге 4 не уточняет названия полей. Вместо этого значение для каждого поля должно предоставляться в том же порядке,
	// в котором они указаны в определении структуры.
	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit.lat, spirit.long)

	// Неважно, как вы инициализируете структуру, вы можете модифицировать специальный символ %v со знаком плюса + для вывода названий полей,
	// как показано в следующем примере. Это особенно полезно для крупных структур.
	fmt.Printf("%+v\n", spirit)
}

// Копирование структур
func task_4() {
	// Когда марсоход достигает востока кратера Брэдбери, бухты Йеллоунайф, место кратера Брэдбери в реальной жизни не меняется,
	// как и в следующем листинге. Переменная curiosity инициализируется с копией значений, находящихся в bradbury,
	// поэтому изменения происходят независимо.
	type location struct {
		lat, long float64
	}

	bradbury := location{-4.5895, 137.4417}
	curiosity := bradbury

	curiosity.lat += 0.0106

	fmt.Println(bradbury, curiosity)

}

// Пример среза структур в Golang\
func task_5() {
	// Лучшим решением станет создание одного среза, где каждое значение является структурой. Тогда каждая локация будет единственной единицей,
	// что можно расширить с названием места высадки или другого необходимого поля, как показано в коде ниже:
	type location struct {
		name      string
		lat, long float64
	}

	locations := []location{
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}

	fmt.Println(locations)
	fmt.Printf("%+v\n", locations)
}

// Кодирование структур в JSON
func task_6() {
	type location struct {
		Lat, Long float64 // Поля должны начинаться с большой буквы
	}

	curiosity := location{-14.5684, 175.472636}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes)) // Выводит: {“Lat”:-4.5895,“Long”:137.4417}

}

// exitOnError выводит любые ошибки и выходит.
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Изменение JSON через теги struct в Golang
func task_7() {
	// Единственное изменение между Листингом 9 и Листингом 10 в использовании тегов struct, что изменяют вывод функции Marshal.
	// Обратите внимание, что поля Lat и Long все еще должны экспортироваться, чтобы пакет json увидел их.
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	curiosyti := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosyti)
	exitOnError(err)

	fmt.Println(string(bytes))
	// Почему поля Lat и Long начинаются с букв в верхнем регистре при кодировании JSON?
	// Пакет json таким образом понимает какие поля нужно экспортировать.
}

// Напишите программу для отображения закодированных данных в JSON трех мест высадки марсохода в Листинге 8.
// JSON должен содержать название каждого места высадки и использовать теги структуры, как показано в Листинге 10.
func task_8() {
	type location struct {
		Name string  `json:"nameLocation"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	mars := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	bytes, err := json.MarshalIndent(mars, "", "  ") // Красивый вывод json
	exitOnError(err)
	JSONBytes(bytes).printJSON()
}

// Методы могут быть привязаны только к типам, определенным в вашем коде (включая пользовательские типы),
// но не к встроенным типам, таким как []byte.
type JSONBytes []byte

func (b JSONBytes) printJSON() {
	fmt.Println(string(b))
}

func main() {
	// task_1() // Объявление структуры в Golang
	// task_2() // Повторное использование структур с типами
	// task_3() // Инициализация структур через композитные литералы
	// task_4() // Копирование структур
	// task_5() // Пример среза структур в Golang
	// task_6() // Кодирование структур в JSON
	// task_7() // Изменение JSON через теги struct в Golang
	task_8() // DZ
}
