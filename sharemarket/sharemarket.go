package main

import (
	"fmt"
	"math"
)

func main() {

	// Max profit from buying and then selling a stock
	price := []int{8, 1, 2, 4, 6, 3}

	max := MaxProfit(price)
	fmt.Println("Max profit for buying and selling a stock of price ", price, " is", max)
}

func MaxProfit(price []int) int {

	// We need to calculate the max profit for selling each day, and then return the max of those
	// The max profit for selling on a given day is the price of that day, minus the min value found up til that day (including that day)

	// Create a slice with the Min values up to a day
	max := math.MinInt

	minUntilNow := math.MaxInt

	for _, v := range price {

		minUntilNow = Min(minUntilNow, v)

		max = Max(max, v-minUntilNow)
	}

	return max
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
