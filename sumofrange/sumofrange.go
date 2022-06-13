package main

import (
	"fmt"
)

type Range struct {
	i, j int
}

func main() {

	// Sum of ranges i - j (inclusive)
	array := []int{8, 1, -2, 4, 6, 3, -5, 0, 8, -4}

	ranges := []Range{
		{0, 0},
		{0, 1},
		{0, 5},
		{0, 9},
		{1, 4},
		{3, 7},
		{6, 9},
		{9, 9},
	}

	sums := SumsOfRanges(array, ranges)
	fmt.Println("Sum of Ranges for this array:  ", array)
	for i, r := range ranges {
		fmt.Printf("Range: %v: Sum: %v\n", r, sums[i])
	}
}

func SumsOfRanges(array []int, ranges []Range) []int {
	// SumOfRange(i, j) = SumUpTo(j) - SumUpTo(i-1)   For i>0

	// Store all SumUpTo(x) values in an array
	sumUpTo := make([]int, len(array))

	sumUpTo[0] = array[0]

	for i := 1; i < len(array); i++ {
		sumUpTo[i] = sumUpTo[i-1] + array[i]
	}

	// now we can iterate over the ranges and return the sums
	sums := make([]int, len(ranges))

	for i, r := range ranges {
		if r.i == 0 {
			sums[i] = sumUpTo[r.j]
		} else {
			sums[i] = sumUpTo[r.j] - sumUpTo[r.i-1]
		}
	}

	return sums
}
