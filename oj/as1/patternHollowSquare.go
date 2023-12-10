package main

import (
	"fmt"
)

/**
print pattern
	* * * *
	*     *
	*     *
	* * * *
*/

func mainPatternHollowSquare() {
	N := 4

	// Loop for each row
	for i := 0; i < N; i++ {
		// Loop for each column in the row
		for j := 0; j < N; j++ {
			// Check if we're on the border, or first/last line
			if i == 0 || i == N-1 || j == 0 || j == N-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		// Newline after each row
		fmt.Println()
	}

}
