package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/blowfish"
)

func main() {
	input := "H5amM2Dq6EAf"
	salt := []byte("N@&vC9X!lPqp38VL")
	blowfishKey := []byte("weborama_encrypt_key") // Замените на тестируемый ключ
	target := "Ct5C656QU4A1288sP9zPPO"

	// SHA-256
	hash := sha256.Sum256([]byte(input + string(salt)))

	// Blowfish
	cipher, err := blowfish.NewCipher(blowfishKey) // Замените на NewSaltedCipher, если есть
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var encrypted [16]byte
	cipher.Encrypt(encrypted[:8], hash[:8])
	cipher.Encrypt(encrypted[8:], hash[8:16])

	// Base64 и замена
	result := base64.StdEncoding.EncodeToString(encrypted[:])
	result = strings.NewReplacer(
		"+", "A",
		"/", "8",
		"=", "O",
	).Replace(result)[:22]

	fmt.Println("Target:", target)
	fmt.Println("Output:", result)
	if result == target {
		fmt.Printf("Success! Key: %s\n", blowfishKey)
	} else {
		fmt.Println("Key is incorrect")
	}
}

// package main

// import (
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"fmt"
// 	"strings"

// 	"golang.org/x/crypto/blowfish"
// )

// type WeboramaEncryptor struct {
// 	blowfishKey []byte
// 	salt        []byte
// }

// func NewWeboramaEncryptor() *WeboramaEncryptor {
// 	return &WeboramaEncryptor{
// 		blowfishKey: []byte("weborama_secure_key_2023"), // Замените на реальный ключ из бинарника
// 		salt:        []byte("N@&vC9X!lPqp38VL"),
// 	}
// }

// func (w *WeboramaEncryptor) Encrypt(input string) string {
// 	// 1. Применяем SHA-256 к salted input
// 	hash := sha256.Sum256([]byte(input + string(w.salt)))

// 	// 2. Blowfish шифрование первых 16 байт хеша
// 	cipher, err := blowfish.NewSaltedCipher(w.blowfishKey, w.salt)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Шифруем блоками по 8 байт
// 	var encrypted [16]byte
// 	cipher.Encrypt(encrypted[:8], hash[:8])
// 	cipher.Encrypt(encrypted[8:], hash[8:16])

// 	// 3. Base64 кодирование с заменой символов
// 	result := base64.StdEncoding.EncodeToString(encrypted[:])
// 	return strings.NewReplacer(
// 		"+", "A",
// 		"/", "8",
// 		"=", "O",
// 	).Replace(result)[:22] // Обрезаем до 22 символов
// }

// func main() {
// 	encryptor := NewWeboramaEncryptor()
// 	input := "H5amM2Dq6EAf"

// 	fmt.Println("Input:   ", input)
// 	fmt.Println("Output:  ", encryptor.Encrypt(input))
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// 	"weboencryptor/encrypter"
// )

// func main() {
// 	enc := encrypter.NewEncripter()
// 	res := CreateWeboId("webo_id.csv", *enc)

// 	newFile, err := os.Create("webo_id_new.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer newFile.Close()

// 	for _, v := range res {
// 		fmt.Fprintf(newFile, "%v,%v\n", v.DecWeboId, v.EncWeboId)
// 	}
// }

// type WeboId struct {
// 	DecWeboId string
// 	EncWeboId string
// }

// func CreateWeboId(fileName string, enc encrypter.Encrypter) []WeboId {
// 	resRead := ReadFile(fileName)
// 	file := strings.Split(strings.ReplaceAll(string(resRead), " ", ""), "\n")

// 	var weboId []WeboId

// 	for _, v := range file {
// 		webo := strings.SplitN(v, ",", 2)
// 		WeboEnc := enc.WeboramaHash(webo[1])
// 		weboId = append(weboId, WeboId{webo[1], string(WeboEnc)})
// 	}

// 	return weboId
// }

// func ReadFile(fileName string) []byte {
// 	file, err := os.ReadFile(fileName)
// 	if err != nil {
// 		panic("Не возможно открыть фаил!")
// 	}

// 	return file
// }
