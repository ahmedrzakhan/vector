package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func findComplement(num int) int {
	// To find complement, we need a mask which has all 1's in binary representation
	// with the same length as num.
	mask := ^0 // All 1's in binary.
	for mask&num > 0 {
		mask <<= 1
	}
	// XOR with mask flips the bits, as 1^1=0 and 0^1=1.
	return ^mask & ^num
}

func main() {
	input := `3
5
1
2147483648`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		complement := findComplement(n)
		fmt.Println(complement)
	}
}
