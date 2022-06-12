package main

import "testing"

func TestCalculateRodMax(t *testing.T) {
	cases := []struct {
		cost     []int
		expected int
	}{
		{[]int{0, 2, 4, 5, 7}, 8},
		{[]int{0, 2}, 2},
	}

	for _, c := range cases {
		max := MaxProfit(c.cost)
		if max != c.expected {
			t.Errorf("Max profit of rod (%v) == %v, but %v was expected", c.cost, max, c.expected)
		}
	}
}
