package main

import "fmt"

/**
Given Array and N, find no of elements larger than their neighbours
and return count
*/

func countLargerThanNeighbour(arr []int) int {
	count := 0
	if len(arr) > 1 { // Ensure there are at least two elements for comparison
		// Check the first element
		if arr[0] > arr[1] {
			count++
		}
		// Check the middle elements
		for i := 1; i < len(arr)-1; i++ {
			if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				count++
			}
		}
		// Check the last element
		if arr[len(arr)-1] > arr[len(arr)-2] {
			count++
		}
	}
	return count
}

func maincountLargerThanNeighbour() {
	arr := []int{1, 2, 3, 4, 2, 1, 6, 5}

	n := len(arr)

	fmt.Println("Number of elements:", n)

	result := countLargerThanNeighbour(arr)
	fmt.Println("res", result)
}
