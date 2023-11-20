package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func main() {
	n := 4
	if isPowerOfTwo(n) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
