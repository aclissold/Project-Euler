package main

import (
	"fmt"
)

func main() {
	var isPrime bool
	i := 10001
	for num := 2; i > 0; num++ {
		isPrime = true
		for j := 2; j < num; j++ {
			if num%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			fmt.Printf("%d prime number #%d\n", num, 10002-i)
			i--
		}
	}
}
