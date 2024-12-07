package main

import (
	"fmt"
	"net/url"
	"os"
)

// В стандартной библиотеке Go есть функция для парсинга веб адресов (см. golang.org/pkg/net/url/#Parse).
// Отобразите ошибку, которая возникает, когда url.Parse используется для неправильного веб адреса вроде того, что содержит пробелы: https://a b.com/.

// Используйте специальный символ %#v с Printf для изучения ошибки. Затем выполните утверждение типа *url.Error для получения доступа и вывода полей
// базовой структуры.

func main() {
	u, err := url.Parse("https://a b.com/")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%#v\n", err)
		// fmt.Printf(url.Error())
		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op:", e.Op)
			fmt.Println("URL:", e.URL)
			fmt.Println("Err:", e.Err)
		}
		os.Exit(1)
	}

	// Парсинг URL
	fmt.Printf("URL: %v\nHost: %v\nForceQuery: %v\nPath: %v\nScheme: %v\nUser: %v\n", u, u.Host, u.ForceQuery, u.Path, u.Scheme, u.User)
}
