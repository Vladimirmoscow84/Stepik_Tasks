package task941

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Task941() {
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
	answer := []string{}
	for _, v := range message {
		if unicode.IsUpper(v) {
			message = strings.ToUpper(message)
			break
		}
	}

	for _, v := range message {
		answer = append(answer, string(v))
	}
	return strings.Join(answer, " ")

}
