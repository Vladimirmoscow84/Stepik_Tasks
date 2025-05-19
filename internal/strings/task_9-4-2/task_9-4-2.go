package task942

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Task942() {
	message := ReadInput()
	result := processMessage(message)
	fmt.Println(result)
}

func ReadInput() string {
	var message string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	message = input
	return message
}
func processMessage(message string) string {
	sl1 := []rune{}
	sl2 := []rune{}
	for _, v := range message {
		if unicode.IsUpper(v) {
			sl1 = append(sl1, v)
		} else {
			sl2 = append(sl2, v)
		}
	}
	return string(append(sl1, sl2...))
}
