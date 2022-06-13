package main

import "testing"

func TestCalculateMaxProfit(t *testing.T) {
	cases := []struct {
		price    []int
		expected int
	}{
		{[]int{8, 1, 2, 4, 6, 3}, 5},
		{[]int{0, 2}, 2},
		{[]int{9, 8, 6, 5, 2}, 0},
	}

	for _, c := range cases {
		max := MaxProfit(c.price)
		if max != c.expected {
			t.Errorf("Max profit of buying and then selling (%v) == %v, but %v was expected", c.price, max, c.expected)
		}
	}
}
