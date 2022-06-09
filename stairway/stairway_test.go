package main

import "testing"

func TestCalculateWays(t *testing.T) {
	cases := []struct {
		in, expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{15, 987},
	}

	for _, c := range cases {
		fn, _ := CalculateWays(c.in)
		if fn != c.expected {
			t.Errorf("Ways(%d) == %d, but %d was expected", c.in, fn, c.expected)
		}
	}
}
