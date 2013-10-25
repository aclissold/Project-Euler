// Solves http://projecteuler.net/problem=439 via a dynamic programming approach.

package main

import (
	"fmt"
	"math"
	"time"
)

// divisorSums holds the previously-computed values of d(k) (a map is basically a Python dictionary)
var divisorSums = make(map[int]int)

func main() {
	//fmt.Println(S(3))
	fmt.Println(S(int(math.Pow(10, 11))) % int(math.Pow(10, 9)))
}

func S(N int) int {
	sum := 0
	divisorSum := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			// startTime := time.Now()
			if computedSum, ok := divisorSums[i*j]; ok { // if computedSum exists at index i*j in the map
				sum += computedSum
			} else {
				// compute divisorSum for i * j and add it to the map
				divisorSum = d(i * j)
				divisorSums[i*j] = divisorSum
				sum += divisorSum
			}
			// fmt.Println("Iteration", j, "after", time.Since(startTime))
		}
	}
	return sum
}

var newSum int
var increment int

func d(k int) int {
	newSum = 0
	// Avoid checking even divisors for odd values of k.
	// I think this cuts out 1/4 of the computations d will ever have to do, but will not actually
	// change the asymptotic behavior of the algorithm.
	if k%2 == 0 {
		increment = 1
	} else {
		increment = 2
	}
	for num := 1; num <= k; num += increment {
		if k%num == 0 {
			newSum += num
		}
	}
	return newSum
}
