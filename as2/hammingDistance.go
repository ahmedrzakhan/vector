package main

import (
	"fmt"
)

/**
Alex has two integers x and y which he received from his two friends.
Find out the hamming distance so that he can know how much different are
these two integers from each other
*/
// hammingDistance calculates the Hamming distance between two integers.
func hammingDistance(x, y int) int {
	distance := 0
	val := x ^ y // XOR x and y, bits will be 1 where x and y differ

	// Count the number of bits set to 1
	for val > 0 {
		distance = distance + val&1 // Increment distance if the least significant bit is 1
		val >>= 1                   // Shift val to the right by one to check the next bit
	}

	return distance
}

func mainHamming() {
	// 15 7 -> 1
	// 0 3 -> 2

	fmt.Println(hammingDistance(15, 7))
}
