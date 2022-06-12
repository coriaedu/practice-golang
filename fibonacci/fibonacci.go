package main

import "fmt"

func main() {

	// Number to calculate its Fibonacci
	n := 15

	fn, fa := CalculateFibonacci(n)
	fmt.Println("Fibonacci for", n, "is", fn, " - Full array:", fa)
}

func CalculateFibonacci(n int) (int, []int) {

	if n < 0 {
		panic("no negatives please")
	}

	if n == 0 {
		return 0, []int{0}
	}

	if n == 1 {
		return 1, []int{0, 1}
	}

	// Create a slice to hold all the Fibonacci's that we calculate
	fa := make([]int, n+1)

	// Initialize F(0) and F(1)
	fa[0] = 0
	fa[1] = 1

	for i := 2; i <= n; i++ {
		fa[i] = fa[i-1] + fa[i-2]
	}

	return fa[n], fa
}
