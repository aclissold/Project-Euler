// Problem 4: Largest Palindrome Product
//
// http://projecteuler.net/problem=4
//
// A palindromic number reads the same both ways. The largest palindrome made from the product of
// two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
    "fmt"
    "log"
    "strconv"
)

var isPalindrome bool
var largestPalindrome = -1 // sentinel value

func main() {
    for i := 999; i > 100; i-- {
        for j := 999; j > 100; j-- {
            isPalindrome = true // assume palindromic to start out
            product := strconv.Itoa(i*j)
            // Iterate inwards on the string representation of the product to check for
            // ...palindromicity...?
            for k := 0; k < len(product)/2; k++ {
                if product[k] != product[len(product)-k-1] {
                    isPalindrome = false
                    break
                }
            }

            if isPalindrome {
                // Update the largest palindrome if necessary
                palindrome, err := strconv.Atoi(product)
                if err != nil {
                    log.Fatal(err)
                }
                if palindrome > largestPalindrome {
                    largestPalindrome = palindrome
                }
            }
        }
    }
    fmt.Println("Largest palindrome:", largestPalindrome)
}
