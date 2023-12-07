package main

import (
	"fmt"
)

/**
Given a non-negative integer n, find out the minimum number of elements in which n can be
broken into such that each element is a power of 2, and the sum of all the elements result
in the integer n.
*/

func countOnes(n int64) int {
	count := 0
	for n > 0 {
		count += int(n & 1) // Increment count if the least significant bit is 1
		n >>= 1             // Shift n to the right by one
	}
	return count
}

func mainCoun() {
	inputs := []int64{1, 15, 2147483648}
	for _, n := range inputs {
		fmt.Printf("Number of elements for %d: %d\n", n, countOnes(n))
	}
}
