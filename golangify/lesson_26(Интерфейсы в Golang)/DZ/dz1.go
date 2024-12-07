package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Удовлетворение требований интерфейса в Go
type location struct {
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

func (l location) String() string {
	return fmt.Sprintf("Переопределение: %v, %v", l.Lat, l.Long)
}

// Напишите метод String для типа coordinate и location и используйте его для отображения координат в более читабельном формате.
type coordinate struct {
	d, m, s float64
	h       rune
}

type Stringer interface {
	String() string
}

func (c coordinate) String() string {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	res := sign * (c.d + c.m/60 + c.s/3600)
	return fmt.Sprintf("%v", res)
}

func errJson(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// decimal конвертирует координаты d/m/s в десятичные градусы
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DD  float64 `json:"decimal"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD:  c.decimal(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

func main() {
	curiosiity := location{
		Lat:  coordinate{4, 35, 22.2, 'S'},
		Long: coordinate{137, 26, 30.12, 'E'},
	}
	fmt.Println(curiosiity)
	c := coordinate{4, 35, 22.2, 'S'}

	bytes, err := json.MarshalIndent(c, "", "  ")
	errJson(err)
	fmt.Println(string(bytes))

	bytes2, err := json.MarshalIndent(curiosiity, "", "  ")
	errJson(err)
	fmt.Println(string(bytes2))
}
