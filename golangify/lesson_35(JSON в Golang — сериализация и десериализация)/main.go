package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Кодирование в формат JSON в Golang
func task_1() {
	u1 := User{1, "Dima", "programist"}
	jsonData, err := json.Marshal(u1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	u2 := []User{
		{2, "Ira", "decret"},
		{3, "Pavel", "manager"},
		{4, "Platon", "baby"},
	}

	jsonData2, err := json.Marshal(u2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData2))

	file, err := os.Create("test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Fprint(file, "[\n")
	for i := range u2 {
		jsonData, err := json.Marshal(u2[i])
		if err != nil {
			log.Fatal(err)
		}
		if i > 0 {
			fmt.Fprint(file, ",\n")
		}
		fmt.Fprintf(file, "%v", string(jsonData))
	}
	fmt.Fprint(file, "\n]")
}

type User struct {
	Id   int
	Name string
	Jobs string
}

// Декодирование данных из JSON файла в Golang
func task_2_my_decodeFile() {
	file, err := os.Open("test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var users []User
	err = decoder.Decode(&users)
	if err != nil {
		log.Fatal(err)
	}

	for i := range users {
		fmt.Println(users[i])
	}
}

// Декодирование данных из JSON в Golang
func task_2() {
	var u1 User

	data := []byte(`{
		"Id" : 1,
        "Name": "John Doe",
        "Occupation": "gardener"
	}`)

	err := json.Unmarshal(data, &u1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u1)
	fmt.Printf("Id: %v, Name: %v, Jobs: %v\n", u1.Id, u1.Name, u1.Jobs)

	var u2 []User

	data2 := []byte(`
    [
        {"Id":2,"Name":"Roger Roe","Occupation":"driver"},
        {"Id":3,"Name":"Lucy Smith","Occupation":"teacher"},
        {"Id":4,"Name":"David Brown","Occupation":"programmer"}
    ]`)

	err2 := json.Unmarshal(data2, &u2)
	if err2 != nil {
		log.Fatal(err2)
	}

	for i := range u2 {
		fmt.Println(u2[i])
	}
}

// Аккуратный вывод формата JSON в Golang
func task_3() {
	birds := map[string]interface{}{
		"sounds": map[string]string{
			"pigeon":  "coo",
			"eagle":   "squak",
			"owl":     "hoot",
			"duck":    "quack",
			"cuckoo":  "ku-ku",
			"raven":   "cruck-cruck",
			"chicken": "cluck",
			"rooster": "cock-a-doodle-do",
		},
	}

	data, err := json.MarshalIndent(birds, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

// Открываем JSON файл в Golang
func task_4() {
	file, err := os.Open("test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var u2 []User

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err3 := json.Unmarshal(fileData, &u2)
	if err3 != nil {
		log.Fatal(err)
	}

	for i := range u2 {
		fmt.Println(u2[i])
	}
	fmt.Println(u2)
}

// Загружаем JSON по ссылке используя HttpClient
func task_5() {
	url := "http://api.open-notify.org/astros.json"

	var newClient = http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := newClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	astros := people{}

	errDecode := json.Unmarshal(data, &astros)
	if errDecode != nil {
		log.Fatal(errDecode)
	}

	for i := range astros.People {
		fmt.Println(astros.People[i])
	}
	fmt.Println(astros)
}

type Astranaut struct {
	Name  string
	Craft string
}

type people struct {
	Number  int
	People  []Astranaut
	Message string
}

func main() {
	// task_1() // Кодирование в формат JSON в Golang
	// task_2_my_decodeFile()
	// task_2_my_decodeFile_part2()
	// task_2() // Декодирование данных из JSON в Golang
	// task_3() // Аккуратный вывод формата JSON в Golang
	// task_4() // Открываем JSON файл в Golang
	task_5() // Загружаем JSON по ссылке используя HttpClient
}
