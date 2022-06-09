package main

import "fmt"

func main() {

	// Number of steps to calculate how many ways to climb it in either 2 or 1 step
	n := 15

	fn, fa := CalculateWays(n)
	fmt.Println("Ways to climb a stair with", n, "steps is", fn, " - Full array:", fa)
}

func CalculateWays(n int) (int, []int) {

	if n < 0 {
		panic("no negatives please")
	}

	if n == 0 {
		return 1, []int{1}
	}

	if n == 1 {
		return 1, []int{1, 1}
	}

	// Create a slice to hold all the ways that we calculate
	fa := make([]int, n+1)

	// Initialize F(0) and F(1)
	fa[0] = 1
	fa[1] = 1

	for i := 2; i <= n; i++ {
		fa[i] = fa[i-1] + fa[i-2]
	}

	return fa[n], fa
}
