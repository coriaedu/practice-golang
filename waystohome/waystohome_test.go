package main

import "testing"

func TestCalculateWays(t *testing.T) {
	cases := []struct {
		cols, rows, expected int
	}{
		{1, 1, 1},
		{15, 1, 1},
		{1, 22, 1},
		{3, 3, 6},
		{5, 3, 15},
		{3, 5, 15},
	}

	for _, c := range cases {
		fn, _ := CalculateWays(c.cols, c.rows)
		if fn != c.expected {
			t.Errorf("Ways(%dx%d) == %d, but %d was expected", c.cols, c.rows, fn, c.expected)
		}
	}
}
