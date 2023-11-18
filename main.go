package main

import (
	"fmt"
	"strings"
)

/**
print pattern
	* * * *
	*     *
	*     *
	* * * *
*/

func main() {
	// var N int
	// fmt.Scan(&N)
	N := 4

	// // Always print the top line
	fmt.Println(strings.Repeat("* ", N))

	for i := 1; i < N-1; i++ {
		// Print leading star
		fmt.Print("*")
		// Print spaces
		for j := 1; j < N-1; j++ {
			fmt.Print("  ")
		}
		// Print trailing star with newline
		fmt.Println(" *")
	}

	// If N is not 1, print the bottom line
	if N != 1 {
		fmt.Println(strings.Repeat("* ", N))
	}

}
