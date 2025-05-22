package task819

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func gravitySwitch(grid [][]int) [][]int {
	slice := make([][]int, len(grid))
	for i := range slice {
		slice[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				for n := 0; n < len(slice); {
					if slice[len(slice)-1-n][j] == 0 {
						slice[len(slice)-1-n][j] = 1
						break
					} else {
						n++
					}
				}
			}
		}
	}

	return slice
}

func Task819() {
	grid := ReadInput()
	result, _ := json.Marshal(gravitySwitch(grid))
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
