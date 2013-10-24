package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(1)
	//sum := 0
	for i := 1; i < 20; i++ {
		fmt.Println(a)
		fmt.Println(b)
		a = b
		b.Add(a, b)
		if i%10000 == 0 {
			fmt.Println("Iteration", i)
		}
	}
}
