package main

import "fmt"

func main() {

	// Number of ways to get home in an m x n grid starting on the top left
	// and only moving down or right.
	cols := 15
	rows := 1

	fnm, fa := CalculateWays(cols, rows)
	fmt.Printf("Ways to climb get home in an %dx%d grid is %d\n", cols, rows, fnm)
	fmt.Println("Full Map:")
	for _, row := range fa {
		fmt.Println(row)
	}
}

func CalculateWays(cols, rows int) (int, [][]int) {

	if cols < 0 || rows < 0 {
		panic("no impossible grids please")
	}

	// Create the array and initialize
	fa := make([][]int, rows+1)
	for i := range fa {
		fa[i] = make([]int, cols+1)
	}

	// Initialize for m = 1 and n =1
	for i := 1; i <= rows; i++ {
		fa[i][1] = 1
	}

	for i := 2; i <= cols; i++ {
		fa[1][i] = 1
	}

	if cols == 1 || rows == 1 {
		return fa[rows][cols], fa
	}

	for i := 2; i <= rows; i++ {
		for j := 2; j <= cols; j++ {
			fa[i][j] = fa[i-1][j] + fa[i][j-1]
		}
	}

	return fa[rows][cols], fa
}
