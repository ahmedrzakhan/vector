package main

import (
	"fmt"
)

/**
if n is power of 2 return True else False without using inbuilt function
*/

/**
if n is power of 2 its only 1 bit would be set ( ... 2**4, 2**3, 2**2, 2**1)

If n is not a power of two, then n has more than one 1 in its binary representation, and
there will be at least one position where both n and n - 1 have a 1. Thus, the AND operation will not result in 0

4 in binary: 100
3 in binary: 011
4 & 3 in binary: 000

If n is not a power of two, then n has more than one 1 in its binary representation, and there will be at least one
position where both n and n - 1 have a 1. Thus, the AND operation will not result in 0.

10 in binary: 1010
9 in binary: 1001
10 & 9 in binary: 1000
*/

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func mainIsPowerOfTwo() {
	n := 4
	if isPowerOfTwo(n) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
