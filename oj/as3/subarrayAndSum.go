package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Given an array of integers of length n and a positive integer K, the task is to find the
count of the longest possible subarrays with the sum of its elements not divisible by K.

Input
4 3
2 3 4 6

Output
1
*/

// O(N^2)
func getCountB(arr []int, N, K int) int {
	maxLength, count := 0, 0
	for i := 0; i < N; i++ {
		sum := 0
		for j := i; j < N; j++ {
			sum = sum + arr[j]
			if sum%K != 0 {
				curLength := j - i + 1
				if curLength > maxLength {
					maxLength = curLength
					count = 1
				} else if curLength == maxLength {
					count++
				}
			}
		}
	}
	return count
}

func getCount(arr []int, N, K int) int {
	// Create a map to store the earliest index where each sum%K value is seen.
	modMap := make(map[int]int)
	modMap[0] = -1 // NOTE: // Initialize with sum 0 at index -1
	maxLength := 0
	sum := 0

	// Loop through the array to fill the modMap and find max subarray length
	for i := 0; i < N; i++ {
		sum = sum + arr[i]
		mod := sum % K

		// Adjust negative mods to be positive
		if mod < 0 {
			mod = mod + K
		}

		// Check if we have seen this mod before
		if idx, exists := modMap[mod]; exists {
			if i-idx > maxLength {
				maxLength = i - idx // Update maxLength
			}
		} else {
			modMap[mod] = i // Store the index of the first occurrence of this mod
		}
	}

	// Find count of subarrays with maxLength
	count := 0
	sum = 0
	for i := 0; i < N; i++ {
		sum = sum + arr[i]
		mod := sum % K

		if mod < 0 {
			mod = mod + K
		}

		if i-modMap[mod] == maxLength {
			count++
			modMap[mod] = i

		}
	}
	return count
}

func mainS() {
	input := `4 3
	2 3 4 6`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	NK := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(NK[0])
	K, _ := strconv.Atoi(NK[1])

	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())

	arr := make([]int, N)
	for index, elem := range arrStr {
		arr[index], _ = strconv.Atoi(elem)
	}

	fmt.Println(getCount(arr, N, K))
}
