package main

import (
	"fmt"
)

func complement(n int) int {
	mask := 7
	return (^n) & mask
}

func mainComplement() {
	n := 5
	fmt.Println(complement(n))
}
