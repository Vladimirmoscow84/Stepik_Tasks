package task82Pt26

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func Task82Pt26() {
	grid := ReadInput()
	result, _ := json.Marshal(sumSquares(grid))
	fmt.Println(string(result))
}

func ReadInput() [][]int {
	var grid [][]int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &grid)
	return grid
}

func sumSquares(grid [][]int) []int {
	arrSum := []int{}
	var sum int
	length := len(grid) / 2
	for i := 0; i < length; i++ {
		for j := i; j < len(grid)-i; j++ {
			for k := i; k < len(grid)-i; k++ {
				if j == i || k == i || j == len(grid)-1-i || k == len(grid)-1-i {
					sum += grid[j][k]
				}
			}

		}
		arrSum = append(arrSum, sum)
		sum = 0
	}
	if len(grid)%2 != 0 {
		arrSum = append(arrSum, grid[length][length])
	}
	return arrSum

}
