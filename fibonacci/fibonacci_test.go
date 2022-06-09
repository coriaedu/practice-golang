package main

import "testing"

func TestCalculateFibonacci(t *testing.T) {
	cases := []struct {
		in, expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{15, 610},
	}

	for _, c := range cases {
		fn, _ := CalculateFibonacci(c.in)
		if fn != c.expected {
			t.Errorf("Fibonacci(%d) == %d, but %d was expected", c.in, fn, c.expected)
		}
	}
}
