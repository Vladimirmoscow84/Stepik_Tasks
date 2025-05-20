package task943

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func Task943() {
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
	sl1 := []string{}
	sl2 := []string{}
	for _, v := range message {
		if unicode.IsUpper(v) {
			sl1 = append(sl1, string(v))
		} else {
			sl2 = append(sl2, string(v))
		}
		sort.Strings(sl1)
	}

	return strings.Join(append(sl1, sl2...), "")
}
