package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
string has 0s and 1s, find maximum lenght od substring (consecutive subsequence) consisting of equal letters
*/

// TC: O(N) SC: O(1)
// Function to calculate the beauty of a string
func calculateBeauty(s string) int {
	maxBeauty, count := 1, 1

	// Iterate through the string and count consecutive characters
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			if count > maxBeauty {
				maxBeauty = count
			}
		} else {
			count = 1
		}
	}

	return maxBeauty
}

// TC O(N * 2^K) SC O(N * K)
// Brute force approach to change up to K characters and calculate the maximum beauty
func calculateMaxBeauty(s string, K int) int {
	// Base case: if K is 0, return the beauty of the string as is
	if K == 0 {
		return calculateBeauty(s)
	}
	// This call is made to establish a baseline for maxBeauty before any changes are made to the string.
	maxBeauty := calculateBeauty(s)

	// Convert string to a slice of runes for easier manipulation
	runes := []rune(s)

	// Try changing each character and calculate the beauty
	for i := 0; i < len(runes); i++ {
		original := runes[i]

		// Change the current character to '0' and '1' and calculate the beauty
		for _, char := range []rune{'0', '1'} {
			if runes[i] != char {
				runes[i] = char
				fmt.Println("rec", runes, i, runes[i], char)
				currentBeauty := calculateMaxBeauty(string(runes), K-1)
				fmt.Println("ds", K, currentBeauty)

				if currentBeauty > maxBeauty {
					maxBeauty = currentBeauty
				}
			}
		}

		// Revert the change
		runes[i] = original
	}

	return maxBeauty
}

// TC - O(1) SC- O(N)
func calculateMaxBeautyOptimized(s string, K int) int {
	maxBeauty, countZero, countOne := 0, 0, 0

	for left, right := 0, 0; right < len(s); right++ {
		// Count the number of 0's and 1's in the current window
		if s[right] == '0' {
			countZero++
		} else {
			countOne++
		}

		// While the window is invalid, shrink it from the left
		for min(countZero, countOne) > K {
			if s[left] == '0' {
				countZero--
			} else {
				countOne--
			}
			left++
		}

		// Update maxBeauty if the current window is larger
		if right-left+1 > maxBeauty {
			maxBeauty = right - left + 1
		}
	}

	return maxBeauty
}

func mainAlice() {
	// var K int
	// var s string

	// Sample Input
	// K = 2
	// s = "0110"

	// Read N, K and the string s if you want to read from stdin
	// fmt.Scan(&N, &K)
	// fmt.Scan(&s)

	// Calculate and print the maximum beauty

	input := `4 2
0110`

	// scanner := bufio.NewScanner(os.Stdin)
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)

	scanner.Scan()
	NK := strings.Fields(scanner.Text())
	K, _ := strconv.Atoi(NK[1])

	scanner.Scan()
	fmt.Println(calculateMaxBeauty(scanner.Text(), K))
	fmt.Println(calculateMaxBeautyOptimized(scanner.Text(), K))
}
