// Solves http://projecteuler.net/problem=439 with a brute force algorithm (T(n) = O(n^3)).

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(S(int(math.Pow(10, 11))) % int(math.Pow(10, 9)))
}

func S(N int) int {
	sum := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			sum += d(i * j)
		}
	}
	return sum
}

func d(k int) int {
	sum := 0
	for num := 1; num <= k; num++ {
		if k%num == 0 {
			sum += num
		}
	}
	return sum
}
