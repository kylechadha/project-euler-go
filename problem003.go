package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	f := primeFactors(600851475143)
	fmt.Println(f)

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
