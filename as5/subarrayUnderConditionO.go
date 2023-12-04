package main

import (
	"fmt"
)

/**
Given an array A, print starting and ending index of all subarrays in the array which has sum 0. (0-indexing)
*/

// Function to find all subarrays with a sum of 0
// TC O(N) SC O(N)
func findSubarraysWithZeroSum(arr []int) [][2]int {
	subarrays := make([][2]int, 0)
	sumMap := make(map[int][]int)

	currentSum := 0

	for i, value := range arr {
		currentSum = currentSum + value

		// Check if currentSum is 0, meaning we found a subarray starting from index 0
		if currentSum == 0 {
			subarrays = append(subarrays, [2]int{0, i})
		}

		// Check if currentSum has been seen before
		if indices, exists := sumMap[currentSum]; exists {
			for _, startIdx := range indices {
				subarrays = append(subarrays, [2]int{startIdx + 1, i})
			}
		}

		// Add the current index to sumMap for the currentSum
		sumMap[currentSum] = append(sumMap[currentSum], i)
	}

	return subarrays
}

func main() {
	arr := []int{6, 3, -1, -3, 4, -2, 2, 4, 6}
	subarrays := findSubarraysWithZeroSum(arr)

	// Sorting the subarrays by their starting index
	for i := 0; i < len(subarrays); i++ {
		for j := i + 1; j < len(subarrays); j++ {
			if subarrays[i][0] > subarrays[j][0] || (subarrays[i][0] == subarrays[j][0] && subarrays[i][1] > subarrays[j][1]) {
				subarrays[i], subarrays[j] = subarrays[j], subarrays[i]
			}
		}
	}

	// Print the subarrays
	for _, subarray := range subarrays {
		fmt.Printf("%d %d\n", subarray[0], subarray[1])
	}
}
