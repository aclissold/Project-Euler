// Solves http://projecteuler.net/problem=439 with a brute force algorithm (T(n) = O(n^3)).

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(S(int(math.Pow(10, 11))) % int(math.Pow(10, 9)))
}

// S(N) returns the sum of particular sums of divisors such that
// S(N) = ∑1≤i≤N ∑1≤j≤N d(i·j)
func S(N int) int {
	sum := 0
	// This doubly-nested for loop gives T(n) = Θ(n²) right off the bat...
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			sum += d(i * j)
		}
	}
	return sum
}

// d returns the sum of the divisors of k. For example, d(4) returns 1 + 2 + 4 = 7.
func d(k int) int {
	sum := 0
	// Determine the divisors of k via iteration over all numbers up to k. Quite possibly the
	// slowest way to perform this task. Maybe one of these algorithms would help:
	// http://en.wikipedia.org/wiki/Integer_factorization#Factoring_algorithms
	for num := 1; num <= k; num++ {
		if k%num == 0 {
			sum += num
		}
	}
	return sum
}
