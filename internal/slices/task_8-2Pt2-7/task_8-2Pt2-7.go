package task82pt27

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func Task82Pt27() {
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
	slice := make([]int, 4)
	for i := range grid {
		for j := range grid[i] {
			if i < len(grid)/2 && j < len(grid[i])/2 {
				slice[0] += grid[i][j]
			} else if i < len(grid)/2 && j >= len(grid[i])/2 {
				slice[1] += grid[i][j]
			} else if i >= len(grid)/2 && j < len(grid[i])/2 {
				slice[2] += grid[i][j]
			} else {
				slice[3] += grid[i][j]
			}
		}
	}
	sort.Ints(slice)
	return slice

}
