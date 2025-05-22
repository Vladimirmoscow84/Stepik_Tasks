package task9412

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Task9412() {
	messages := ReadInput()
	result := sortByDigitCount(messages)
	fmt.Println(result)
}

func ReadInput() []string {
	var messages []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &messages)
	return messages

}

func sortByDigitCount(messages []string) string {
	answer := make([]string, 0)
	var digit int
	sl := make([]int, 0)
	mapa := make(map[int][]string)
	for _, v := range messages {
		if strings.Contains(v, "_") {
			sl2 := strings.Split(v, "_")
			digit = len(sl2[1])
		} else {
			digit = 0
		}
		mapa[digit] = append(mapa[digit], v)
	}
	for i := range mapa {
		sl = append(sl, i)
	}
	sort.Ints(sl)
	for _, v := range sl {
		answer = append(answer, mapa[v]...)
	}

	return strings.Join(answer, ", ")
}

// Другое решение
// Устойчивая сортировка

//  func getDigitsCount(s string) int {
// 	 cnt := 0
// 	 for _, l := range s {
// 	 if unicode.IsDigit(l) {
// 	  cnt++
// 	 }
// 	}
// 	return cnt
//    }

//    func sortByDigitCount(messages []string) string {
//    sort.SliceStable(messages, func(i, j int)bool{
// 	return getDigitsCount(messages[i])<getDigitsCount(messages[j])
//    })
//    return strings.Join(messages, ", ")
//    }
