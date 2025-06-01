package task10113

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func Task1013() {
	message1, message2, message3 := ReadInput()
	answer := findLongestCommonWord(message1, message2, message3)
	fmt.Println(answer)

}
func ReadInput() (string, string, string) {
	var message1, message2, message3 string
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	message1 = values[0]
	message2 = values[1]
	message3 = values[2]
	return message1, message2, message3

}

func findLongestCommonWord(message1, message2, message3 string) string {
	length := 0
	answer := ""
	sl := strings.Fields(message1)
	for _, v := range sl {
		lword := utf8.RuneCountInString(v)
		if strings.Contains(message2, v) && strings.Contains(message3, v) && lword > length {
			length = lword
			answer = v
		}
	}
	return fmt.Sprintf("Самое длинное общее слово: %s", answer)
}
