package encrypter

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

type Encrypter struct {
	salt string
}

func NewEncripter() *Encrypter {
	salt := "N@&vC9X!lPqp38VL"
	return &Encrypter{
		salt: salt,
	}
}

func (enc *Encrypter) WeboramaHash(input string) string {
	// 1. Добавляем соль
	data := enc.salt + input

	// 2. Вычисляем SHA-256
	hash := sha256.Sum256([]byte(data))

	// 3. Берём первые 15 байт хеша (не 20!)
	truncated := hash[:15]

	// 4. Кодируем в Base64
	encoded := base64.StdEncoding.EncodeToString(truncated)

	// 5. Заменяем символы чтобы соответствовать Weborama
	result := strings.Map(func(r rune) rune {
		switch {
		case r == '+':
			return 'A'
		case r == '/':
			return '8'
		default:
			return r
		}
	}, encoded)

	return result
}
