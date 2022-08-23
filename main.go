package main

import (
	"fmt"
	"recursion/fibonacci"
)

func main() {
	var len uint = 10

	for i := uint(1); i <= len; i++ {
		fibonacciNumber := fibonacci.GetFibonacciNumber(i)

		fmt.Println(fibonacciNumber)
	}
}
