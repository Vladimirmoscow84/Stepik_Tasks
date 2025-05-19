package task82Pt23

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func Task82Pt23() {

	grid := ReadInput()
	result, _ := json.Marshal(propagateLight(grid))
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

func propagateLight(grid [][]int) [][]int {
	rows, cols := len(grid), len(grid[0])
	if rows < 3 || cols < 3 {
		return grid
	}

	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				for ii := max(r-1, 0); ii < min(r+2, rows); ii++ {
					for jj := max(c-1, 0); jj < min(c+2, cols); jj++ {
						result[ii][jj] = 1
					}
				}
			}
		}
	}

	return result
}
