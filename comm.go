package main

import (
	"bytes"
	"fmt"
)

// вводим текст и функция разделяет его по 3 символа через запятую
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	// делаем записи
	buf.WriteString(s[:n-3])
	buf.WriteString(", ")
	buf.WriteString(s[n-3:])
	// склеиваем и возвращаем записи
	return buf.String()
}

func main() {

	var input string
	fmt.Println("Введите строку")
	fmt.Scanln(&input)
	fmt.Println(comma(input))

}
