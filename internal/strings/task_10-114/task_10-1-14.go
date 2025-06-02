package task10114

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Task10114() {
	message := ReadInput()
	answer := findLongestCombinedWord(message)
	fmt.Println(answer)

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

func findLongestCombinedWord(message string) string {
	sl := strings.Fields(message)
	mess := ""
	//var slice []string
	maxLen := 0
	answer := ""

	for i := 0; i < len(sl); i++ {
		mess = sl[i]
		for j := 0; j < len(sl); j++ {
			if i == j {
				continue
			}
			if strings.Contains(mess, sl[j]) {
				_, a, ok := strings.Cut(mess, sl[j])
				if ok {
					mess = a
				}
				if len(mess) == 0 {
					break
				}
			}
		}
		if len(mess) == 0 {
			if len(sl[i]) > maxLen {
				maxLen = len(sl[i])
				answer = sl[i]
			}
			break
		}

	}

	return answer

}
