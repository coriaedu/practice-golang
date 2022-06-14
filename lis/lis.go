package main

import (
	"fmt"
)

func main() {

	nums := []int{11, 7, 12, 34, 1, 2, 7, 8}

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

	// Store all sequences for each LIS
	seq := make([][]int, len(nums))

	// initialize the first value
	lises[0] = 1
	seq[0] = []int{nums[0]}
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

		if max_j >= 0 {
			seq[i] = append(seq[max_j], nums[i])
		} else {
			seq[i] = []int{nums[i]}
		}
		// fmt.Printf("LIS[%v] = %v\n", i, lises[i])

		if lises[i] > maxLis {
			maxLis = lises[i]
			max_i = i
		}

	}

	return maxLis, seq[max_i]
}
