package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Объявление карты в Golang
func task_1() {
	temperature := map[string]int{
		"Земля": 15,
		"Марс":  -65,
	}

	temperature["Земля"] = 16 // Вносим изменения в карту
	fmt.Printf("Температура земли изменилась с 15 градусов до %v градусов!\n", temperature["Земля"])
	temperature["Венера"] = 464
	fmt.Printf("В карту температур добавлена %v!\n%v\n", "Венера", temperature)

	// При попытке получить доступ к ключу, которого нет в карте, результатом будет нулевое значение типа (int)
	mooon := temperature["Луна"]
	fmt.Println(mooon)

	// temperature["Луна"] = -30

	// В Go используется синтаксис comma, ok, что нужен для обозначения разницы между ключом "Луна",
	// которого нет в карте, и тем, что находится в карте с температурой 0° C:
	if moon, ok := temperature["Луна"]; ok {
		fmt.Printf("Температура Луны равна: %v\n", moon)
	} else {
		fmt.Println("Где Луна?")
	}
}

// Копируются ли карты в Golang?
func task_2() {
	planets := map[string]string{
		"Земля": "Сектор ZZ9",
		"Марс":  "Сектор ZZ9",
	}

	// Ранее мы упоминали, что массивы копируются во время присваивания к новым переменным или передачи к функциям или методам.
	// То же самое верное в отношении примитивных типов вроде int или float64.

	// Карты устроены иначе. В следующем примере planets и planetsMarkII делят одну и ту же базовую информацию.
	// Как видите, изменения в одном, затрагивают и другой. Это не всегда кстати.
	planetsMarkII := planets
	planets["Земля"] = "Упс!"

	fmt.Printf("Оригинальный масив - %v\nКопия - %v\n", planets, planetsMarkII)

	delete(planets, "Земля") // Земля удаляется из карты
	fmt.Printf("Оригинальный масив - %v\nКопия - %v\n", planets, planetsMarkII)
}

// Предварительное обозначение карты через make
func task_3() {
	temperature := make(map[string]int, 8)
	temperature["Земля"] = 16
	fmt.Println(temperature)
}

// Использование карты для подсчета частоты использования элементов
func task_4() {
	temperature := []float64{-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0}

	frequency := make(map[float64]int, 10)

	for _, i := range temperature {
		frequency[i]++
	}

	for k, v := range frequency {
		fmt.Printf("Температура: %+.2f, всречаеться %d раз\n", k, v)
	}
}

// Группирование данных с картами и срезами Golang
func task_5() {
	temperature := []float64{-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0}

	group := make(map[float64][]float64)

	for _, t := range temperature {
		g := math.Trunc(t/10) * 10
		group[g] = append(group[g], t)
	}

	for k, v := range group {
		fmt.Printf("Группа: %+.2f, Значения: %v\n", k, v)
	}
}

// Множества в Golang
func task_6() {
	temperature := []float64{-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0}

	set := make(map[float64]bool) // Обьявление масива 'множество'

	for _, t := range temperature {
		set[t] = true
	}

	if set[32.0] {
		fmt.Printf("Температура: 32.0 входит в множество!\n")
	}

	fmt.Printf("Масив 'множество': %v\n", set)

	// Видно, что карта содержит по одному ключу для каждой температуры, дубликаты удаляются.
	// У ключей карты произвольный порядок, поэтому перед их сортировкой температуры нужно конвертировать обратно в срез:
	unique := make([]float64, 0, len(set))
	for k, _ := range set {
		unique = append(unique, k)
	}

	sort.Float64s(unique)
	fmt.Printf("Отсортированный срез: %v\n", unique)
}

//Напишите функцию для подсчета частоты упоминания слов в строке текста и возвращения карты со словами и числом, указывающем,
// сколько раз они употребляются. Функция должна конвертировать текст в нижний регистр и обрезать знаки препинания. Используйте пакет strings.
// Функции, которые пригодятся для выполнения данного задания: Fields, ToLower и Trim.

// Используйте функцию для подсчета частоты слов следующего отрывка текста и последующего вывода числа употребления каждого слова,
// что встречается более одного раза.
func task_7() {
	str := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."
	strArr := strings.Fields(str)

	group := make(map[string]int)
	for _, i := range strArr {
		group[strings.Trim(i, `,;"-.`)]++
		// fmt.Println(strings.ToLower(strings.Trim(i, ",;.")))
	}
	for k, v := range group {
		if v > 1 {
			fmt.Printf("%v: встречаеться - %v раз\n", k, v)
		}
	}
}

// Решение преподователя
func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	frequency := make(map[string]int, len(words))
	for _, word := range words {
		word = strings.Trim(word, `.,"-`)
		frequency[word]++
	}
	return frequency
}

func task_8() {
	text := `As far as eye could reach he saw nothing but the stems of the
	great plants about him receding in the violet shade, and far overhead
	the multiple transparency of huge leaves filtering the sunshine to the
	solemn splendour of twilight in which he walked. Whenever he felt able
	he ran again; the ground continued soft and springy, covered with the
	same resilient weed which was the first thing his hands had touched in
	Malacandra. Once or twice a small red creature scuttled across his
	path, but otherwise there seemed to be no life stirring in the wood;
	nothing to fear -- except the fact of wandering unprovisioned and alone
	in a forest of unknown vegetation thousands or millions of miles
	beyond the reach or knowledge of man.`

	frequency := countWords(text)
	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("%d %v\n", count, word)
		}
	}
}

func main() {
	// task_1() // Объявление карты в Golang
	// task_2() // Копируются ли карты в Golang?
	// task_3() // Предварительное обозначение карты через make.
	// task_4() // Использование карты для подсчета частоты использования элементов
	// task_5() // Группирование данных с картами и срезами Golang
	// task_6() // Множества в Golang
	// task_7() // DZ
	task_8() // Решение преподователя
}
