package task1013

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Task1013() {
	message := ReadInput()
	result := getTopHashtags(message)
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

func getTopHashtags(message string) string {
	re := regexp.MustCompile(`[А-я]{5,}`)
	sl := re.FindAllString(strings.ToLower(message), -1)
	return "#" + strings.Join(sl[len(sl)-5:], " #")

}
