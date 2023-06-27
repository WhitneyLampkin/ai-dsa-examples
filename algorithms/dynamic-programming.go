/*
In the Fibonacci sequence, each number is the sum of the two preceding ones. The Dynamic Programming approach uses a cache (an array in this case) to store previously calculated Fibonacci numbers, which avoids redundant calculations.

In the example, we define the Fibonacci function that takes an integer n as input and returns the nth Fibonacci number using dynamic programming. We initialize the cache with base cases (cache[0] = 0 and cache[1] = 1), and then we calculate Fibonacci numbers from 2 to n by summing the previous two numbers. Finally, we return the nth Fibonacci number.

When running the program, it will output:

The 10th Fibonacci number is: 55
This demonstrates how dynamic programming can efficiently compute the Fibonacci sequence by avoiding redundant calculations and reusing previously calculated values.
*/

package main

import "fmt"

// Fibonacci calculates the nth Fibonacci number using dynamic programming
func Fibonacci(n int) int {
	// Create a cache to store previously calculated Fibonacci numbers
	cache := make([]int, n+1)

	// Base cases
	cache[0] = 0
	cache[1] = 1

	// Calculate Fibonacci numbers from 2 to n
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
}

func main() {
	n := 10
	fib := Fibonacci(n)
	fmt.Printf("The %dth Fibonacci number is: %d\n", n, fib)
}