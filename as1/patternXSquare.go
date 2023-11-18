package main

import (
	"fmt"
)

func mainPatternXSquare() {
	var N int
	fmt.Scan(&N)

	// Loop for each row
	for i := 0; i < N; i++ {
		// Loop for each column in the row
		for j := 0; j < N; j++ {
			// Print '*' at borders and diagonals, ' ' otherwise
			if i == 0 || j == 0 || i == N-1 || j == N-1 || i == j || i+j == N-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		// Newline after each row
		fmt.Println()
	}
}
