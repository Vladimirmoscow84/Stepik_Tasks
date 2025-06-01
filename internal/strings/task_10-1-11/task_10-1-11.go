package task10111

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Task10111() {
	message := ReadInput()
	answer := getPairs(message)
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
func getPairs(message string) string {
	mapa := make(map[string]int) //подсчет комбинаций соседний символов
	slice := make([]string, 0)   // промежуточный слайс соседних символов
	answer := make([]string, 0)  // слайс для формирования ответа
	runes := []rune(message)
	for i, j := 0, 1; i < len(runes)-1; i, j = i+1, j+1 {
		val := string(runes[i]) + string(runes[j])
		mapa[val]++
		slice = append(slice, val)
	}
	for _, v := range slice {
		if _, ok := mapa[v]; ok {
			answer = append(answer, fmt.Sprintf("%s: %d", v, mapa[v]))
			delete(mapa, v)
		}
	}
	return strings.Join(answer, ", ")

}
