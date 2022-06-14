package main

import "testing"

func TestCalculateSumsOfRanges(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{ //cases
		{
			[]int{11, 23, 10, 37, 21, 50, 80}, 5,
		},
		{
			[]int{11, 8, 5}, 1,
		},
		{
			[]int{11}, 1,
		},
		{
			[]int{1, 2, 3}, 3,
		},
		{
			[]int{11, 7, 12, 34, 1, 2, 7, 8}, 4,
		},
	}

	for _, c := range cases {
		lis, _ := LongestIncreasingSubsequence(c.nums)

		if lis != c.expected {
			t.Errorf("LIS of nums %v, was %v but %v was expected", c.nums, lis, c.expected)
		}
	}
}
