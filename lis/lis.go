package main

import (
	"fmt"
)

func main() {

	nums := []int{11, 23, 10, 37, 21, 50, 80}

	lis, max_seq := LongestIncreasingSubsequence(nums)
	fmt.Println("Longest Increasing Subsequence for these nums", nums, "is", lis, "with sequence", max_seq)
}

func LongestIncreasingSubsequence(nums []int) (int, []int) {
	if len(nums) < 1 {
		return 0, []int{}
	}

	// LIS[i] = 1 + the max value for previous LIS's whose index's value is less than nums[i]

	// Store all LIS(i) values in an array
	lises := make([]int, len(nums))
	// initialize the first value
	lises[0] = 1

	// Store the previous index for each indexe's LIS
	seq := make([]int, len(nums))
	seq[0] = -1 // indicates end of

	//fmt.Printf("LIS[0] = %v\n", lises[0])

	// The global max LIS
	maxLis := 1
	max_i := 0

	for i := 1; i < len(nums); i++ { // Let's get all the LIS[i]

		// Let's get the j's max
		innerMax := 0
		max_j := -1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if lises[j] > innerMax {
					innerMax = lises[j]
					max_j = j
				}
			}
		}

		lises[i] = 1 + innerMax
		seq[i] = max_j

		// fmt.Printf("LIS[%v] = %v\n", i, lises[i])

		if lises[i] > maxLis {
			maxLis = lises[i]
			max_i = i
		}

	}

	return maxLis, sequence(nums, seq, max_i)
}

func sequence(nums []int, seq []int, end int) []int {

	final := []int{}

	for end >= 0 {
		final = append([]int{nums[end]}, final...)
		end = seq[end]
	}

	return final
}
