package task944

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Task944() {
	message := ReadInput()
	result := processMesage(message)
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
func processMesage(message string) string {
	return strings.Title(strings.ToLower(message))

}
