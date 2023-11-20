package main

import (
	"fmt"
)

/**
find element with element with lower index to that  are smaller
and element with higher indexes are higher
*/

// Function to find the equilibrium index
func elementInMiddle(arr []int) int {
	N := len(arr)
	response := -1
	for i := 1; i < N; i++ {
		leftLess := true
		rightLess := true
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[i] {
				leftLess = false
				break
			}

		}

		for k := i + 1; k < N; k++ {
			if arr[i] > arr[k] {
				rightLess = false
				break
			}
		}

		if leftLess && rightLess {
			response = arr[i]
			break
		}

	}
	return response
}

func mainElementInMiddle() {
	arr := []int{4, 3, 6, 7, 8}
	// Find and print the equilibrium index
	fmt.Println("elementInMiddle", elementInMiddle(arr))
}
