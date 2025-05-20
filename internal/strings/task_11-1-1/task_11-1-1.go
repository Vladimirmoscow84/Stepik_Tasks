package task1111

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Task1111() {
	message, shift := ReadInput()
	result := encrypt(message, shift)
	fmt.Println(result)

}
func ReadInput() (string, int) {
	var message string
	var shift int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	message = values[0]
	shift, _ = strconv.Atoi(values[1])
	return message, shift

}

func encrypt(message string, shift int) string {
	var answer string
	var w int
	shift = shift % 26 // убираем зацикливание алфавита, так как в инглише 26 символов
	for _, v := range message {
		if unicode.IsLetter(v) { // проверяем, что символ является буквой
			if unicode.IsUpper(v) { //проверяем, что буква Заглавная
				switch {
				case int(v)+shift <= 90: // проверям выход за границу кодировки заглавных букв
					w = int(v) + shift
					answer += string(rune(w))
				case int(v)+shift > 90:
					w = 64 + (int(v) + shift - 90)
					answer += string(rune(w))
				}
			}
			if unicode.IsLower(v) { //проверяем, что буква строчная
				switch {
				case int(v)+shift <= 122: // проверям выход за границу кодировки заглавных букв
					w = int(v) + shift
					answer += string(rune(w))
				case int(v)+shift > 122:
					w = 96 + (int(v) + shift - 122)
					answer += string(rune(w))
				}
			}
		} else {
			answer += string(v)
		}
	}
	return string(answer)
}
