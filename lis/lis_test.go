package main

import (
	"reflect"
	"testing"
)

func TestCalculateSumsOfRanges(t *testing.T) {
	cases := []struct {
		nums        []int
		expected    int
		expectedSeq []int
	}{ //cases
		{
			[]int{11, 23, 10, 37, 21, 50, 80}, 5, []int{11, 23, 37, 50, 80},
		},
		{
			[]int{11, 8, 5}, 1, []int{11},
		},
		{
			[]int{11}, 1, []int{11},
		},
		{
			[]int{1, 2, 3}, 3, []int{1, 2, 3},
		},
		{
			[]int{11, 7, 12, 34, 1, 2, 7, 8}, 4, []int{1, 2, 7, 8},
		},
	}

	for _, c := range cases {
		lis, seq := LongestIncreasingSubsequence(c.nums)

		if lis != c.expected {
			t.Errorf("LIS of nums %v, was %v but %v was expected", c.nums, lis, c.expected)
		}

		if !reflect.DeepEqual(seq, c.expectedSeq) {
			t.Errorf("SEQ of nums %v, was %v but %v was expected", c.nums, seq, c.expectedSeq)
		}
	}
}
