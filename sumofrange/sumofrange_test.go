package main

import "testing"

func TestCalculateSumsOfRanges(t *testing.T) {
	cases := []struct {
		array    []int
		ranges   []Range
		expected []int
	}{ //case 1
		{
			[]int{8, 1, -2, 4, 6, 3, -5, 0, 8, -4},
			[]Range{
				{0, 0},
				{0, 1},
				{0, 5},
				{0, 9},
				{1, 4},
				{3, 7},
				{6, 9},
				{9, 9},
			},
			// expexted sums
			[]int{8, 9, 20, 19, 9, 8, -1, -4},
		},
		//case 2
		{
			[]int{5},
			[]Range{
				{0, 0},
			},
			// expexted sums
			[]int{5},
		},
	}

	for _, c := range cases {
		sums := SumsOfRanges(c.array, c.ranges)

		for i, sum := range sums {
			if sum != c.expected[i] {
				t.Errorf("Sum of range (%v) for Array %v, was %v but %v was expected", c.ranges[i], c.array, sum, c.expected[i])
			}
		}
	}
}
