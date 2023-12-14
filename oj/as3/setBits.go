package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Given an integer output the count of Set Bits present in it's binary representation.
*/

/**
can also mod with 2 and right shift and counter
*/

// countSetBits returns the count of set bits in the binary representation of n.
func countSetBits(n int) int {
	count := 0
	for n > 0 {
		if int(n&1) == 1 {
			count++ // Increment count if the least significant bit is 1
		}
		n >>= 1 // Shift n to the right by one to check the next bit
	}
	return count
}

func mainNum() {
	input := `4
1
3
7
6`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < T; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		fmt.Println(countSetBits(N))
	}
}
