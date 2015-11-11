package main

import "fmt"

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

func main() {
	m := multiples(10)
	s := sum(m)

	fmt.Println(s)
}

func multiples(n int) []int {
	var multiples []int
	for i := 1; i < n; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			multiples = append(multiples, i)
		}
	}

	return multiples
}

func sum(m []int) int {
	var sum int
	for _, elem := range m {
		sum += elem
	}

	return sum
}
