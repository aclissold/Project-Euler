// Problem 25: 1000-digit Fibonacci Number
//
// http://projecteuler.net/problem=25
//
// The Fibonacci sequence is defined by the recurrence relation:
//
// Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
// Hence the first 12 terms will be:
//
// F1 = 1
// F2 = 1
// F3 = 2
// F4 = 3
// F5 = 5
// F6 = 8
// F7 = 13
// F8 = 21
// F9 = 34
// F10 = 55
// F11 = 89
// F12 = 144
//
// The 12th term, F12, is the first term to contain three digits.
//
// What is the first term in the Fibonacci sequence to contain 1000 digits?

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(0), big.NewInt(1)
    i := 0
	for {
		if len(a.String()) >= 1000 {
			fmt.Printf("Fibonacci term #%d is the first to contain 1000 digits.\n", i)
			break
		} else {
			a, b = b, a.Add(a, b)
            i++
		}
	}
}
