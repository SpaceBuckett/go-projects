package main 

import "fmt"

func main() {
	col := 3
	row := 3
	someRandomNumbers := [][]int{{9, 8, 8}, {1, 3, 4}, {23, 1, 1}}

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			fmt.Printf("%d ", someRandomNumbers[i][j])
		}
		fmt.Println()
	}
}
