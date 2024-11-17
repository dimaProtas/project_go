package main

import "fmt"

func task_1() {
	// К счастью, можно сгруппировать связанные поля вместе с помощью структур и композиции.

	bradbury := Location{-4.5895, 137.4417, World{radius: 3389.5}}
	t := Temperature{high: -1.0, low: -78.0}
	reports := Report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", reports)
	fmt.Printf("Максимальная емпреатура на Bradbury, равна %v\n", reports.temperature.high)

	fmt.Printf("Средняя температура на Bradbury равна: %.2f\n", reports.temperature.Average())
	fmt.Printf("Средняя температура из отчета на Bradbury равна: %.2f\n", reports.AverageReportsTemperature())
}

// Встраивание методов в Golang
func task_2() {
	// Все методы типа Temperature автоматически делаются доступными через тип report:
	report := Report2{
		sol:         15,
		Location:    Location{-4.5895, 137.4417, World{radius: 3389.5}},
		Temperature: Temperature{high: -1.0, low: -78.0},
	}

	fmt.Printf("Средняя температура: %+v\n", report.Average())
	// Хотя название поля не уточняется, поле все еще существует с тем же названием, что и внедренный тип.
	// Получить доступ к полю temperature можно следующим образом:
	fmt.Printf("Средняя температура: %+v\n", report.Temperature.Average())
	// Внедрение не только встраивает методы. Поля внутренней структуры доступны из внешней структуры.
	// В дополнение к report.temperature.high вы можете получить доступ к  температуре с помощью report.high, что делается следующим образом:
	fmt.Printf("Досступ к температуре без уточнения Temperature: %v\n", report.high)

	fmt.Printf("Вычисляем дни на марсе - %v\n", report.sol.Days(1446))
	// Так как Days есть и у типа Location и sol то go не понимает какой метод вызывать
	// Определили приоретет для метода Days в типе Reports (теперь без ошибки)
	fmt.Printf("Вычисляем дни на марсе - %v\n", report.Days(1446)) // Метод закрепляеться за типом или структурой котоорой пренадлежит

	// Будет ли report.lat правильно отображаться? Если да, то к какому полю он отсылается в Листинге 4?
	fmt.Println(report.lat) // Ответ: да. Пренадлежит полю Location.
}

// Столкновение, или коллизия названий в Go
func task_3() {
	report_mars_opportunity := Report2{
		sol:         15,
		Location:    Location{-1.9462, 354.4734, World{radius: 3389.5}},
		Temperature: Temperature{high: -1.0, low: -78.0},
	}

	report_mars_spirit := Report2{
		sol:         15,
		Location:    Location{-14.5684, 175.472636, World{radius: 3389.5}},
		Temperature: Temperature{high: -1.0, low: -78.0},
	}

	fmt.Printf("Растояние от opportunity до spirti равно: %.2f\n", report_mars_opportunity.Location.Days(report_mars_spirit.Location))
}

func main() {
	// task_1() // Композиция структур в Golang
	// task_2() // Встраивание методов в Golang
	task_3() // Столкновение, или коллизия названий в Go
}
