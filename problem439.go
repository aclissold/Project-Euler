// Solves http://projecteuler.net/problem=439 via a dynamic programming approach.

package main

import (
	"fmt"
	"math"
	"time"
)

// divisorSums holds the previously-computed values of d(k)
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
			startTime := time.Now()
			if computedSum, ok := divisorSums[i*j]; ok {
				sum += computedSum
			} else {
				divisorSum = d(i * j)
				divisorSums[i*j] = divisorSum
				sum += divisorSum
			}
			fmt.Println("Iteration", j, "after", time.Since(startTime))
		}
	}
	return sum
}

var newSum int
var increment int

func d(k int) int {
	newSum = 0
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
