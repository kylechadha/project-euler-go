package main

import (
	"fmt"
	"log"
	"time"
)

// Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func main() {
	// Let's do a bit of back-of-the-envelope profiling.
	start := time.Now()

	f := primeFactors(600851475143)
	fmt.Println(f)

	// Solves in 200µs! A more traditional approach (better suited to a larger n) solves at 5-7ms.
	elapsed := time.Since(start)
	log.Printf("Factorization took %s", elapsed)
}

func primeFactors(n int) int {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			n = primeFactors(n / i)
			break
		}
	}
	return n
}
