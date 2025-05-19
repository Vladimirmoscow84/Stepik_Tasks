package task941

import (
	"bufio"
	"fmt"
	"os"
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

}
