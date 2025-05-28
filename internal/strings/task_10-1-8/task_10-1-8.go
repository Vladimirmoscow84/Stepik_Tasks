package task1018

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Task1018() {
	message1, message2 := ReadInput()
	result := countWordOccurrences(message1, message2)
	fmt.Println(result)

}

func ReadInput() (string, string) {
	var message1 string
	var message2 string
	var values []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	message1 = values[0]
	message2 = values[1]
	return message1, message2
}

func countWordOccurrences(message1, message2 string) string {
	answer := ""
	sl := strings.Split(message1, " ")
	for _, v := range sl {
		answer += fmt.Sprintf("%s: %d, ", v, strings.Count(message2, v))
	}
	answer = strings.TrimSuffix(answer, ", ")

	return answer
}
