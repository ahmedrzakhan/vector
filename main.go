package main

import (
	"fmt"
)

// Function to find the equilibrium index
func findEquilibrium(arr []int) int {
	totalSum := 0
	leftSum := 0

	// Calculate the total sum of the array
	for _, num := range arr {
		totalSum += num
	}

	// Iterate through the array to find the equilibrium index
	for i, num := range arr {
		totalSum -= num // Now totalSum is right sum for index i

		if leftSum == totalSum {
			return i // Equilibrium index found
		}

		leftSum += num // Update leftSum for next iteration
	}

	return -1 // No equilibrium index found
}

func main() {

	// Read the size of the array

	arr := []int{15, 1, 5, 5, 5}
	// Find and print the equilibrium index
	fmt.Println(findEquilibrium(arr))
}
