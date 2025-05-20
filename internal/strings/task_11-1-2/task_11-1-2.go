package task1112

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task1112() {
	message, shift := ReadInput()
	result := decrypt(message, shift)
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

func decrypt(message string, shift int) string {
	var answer string
	for _, v := range message {
		if v > 64 && v < 91 { // проверка, что буква заглавная
			v = 90 - v
			v = 90 - rune((int(v)+shift)%26)
		} else if v > 96 && v < 123 { // проверка, что буква строчная
			v = 122 - v
			v = 122 - rune((int(v)+shift)%26)
		}

		answer += string(v)
	}
	return answer
}
