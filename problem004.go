package main

import (
	"fmt"
	"strconv"
)

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func main() {
	largest := palindromicNum(999, 999)

	fmt.Println(largest)
}

// Finds the largest palindromic number made from the product of a and b.
func palindromicNum(x, y int) string {
	for a := x; a > 0; a-- {
		for b := y; b > 0; b-- {
			// Calculate the product and determine the midpoint.
			p := strconv.Itoa(a * b)
			mid := len(p) / 2

			// Test palindromity (its a real word, I swear).
			if p[:mid] == reverse(p[mid:]) {
				return p
			}
		}
	}

	return "No solution found."
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
