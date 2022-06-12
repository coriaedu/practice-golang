package main

import (
	"fmt"
	"math"
)

func main() {

	// Max value from selling a rod
	cost := []int{0, 2, 4, 5, 7}

	max := MaxProfit(cost)
	fmt.Println("Max profit of cutting a rod of costs ", cost, "steps is", max)
}

func MaxProfit(cost []int) int {

	n := len(cost) - 1
	if n < 0 {
		panic("real rods please")
	}

	// Initialize the slice to store the rod[x] values
	rod := make([]int, n+1)
	rod[0] = cost[0]

	// rod[i] = max (cost[j] + rod[i-j]) for all 0 > j >= i

	for i := 1; i <= n; i++ { // this will calculate rod[i]
		max := math.MinInt // initialize the max for this "i"
		// For this i, we need to calculate with the different costs[j]
		for j := 1; j <= i; j++ {
			max = Max(max, cost[j]+rod[i-j])
		}
		rod[i] = max
	}

	return rod[n]
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
