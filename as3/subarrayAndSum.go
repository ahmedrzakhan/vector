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
// TC - O(N^2) SC - O(1)
// func countSubarrayAndSumBrute(arr []int, N int, K int) int {
// 	longestLength := 0
// 	count := 0

// 	// Brute force approach to check all subarrays
// 	for i := 0; i < N; i++ {
// 		sum := 0
// 		for j := i; j < N; j++ {
// 			sum += arr[j]
// 			if sum%K != 0 {
// 				length := j - i + 1
// 				if length > longestLength {
// 					longestLength = length
// 					count = 1 // reset count for a new longest length
// 				} else if length == longestLength {
// 					count++
// 				}
// 			}
// 		}
// 	}

// 	return count
// }

// TC - O(N) SC - O(N)
func countSubarrayAndSum(arr []int, N int, K int) int {
	prefixMod := make(map[int]int)
	longestLength := 0
	count := 0
	sum := 0

	// Initialize the hashmap for the case when subarray starts from index 0
	prefixMod[0] = -1

	for i := 0; i < N; i++ {
		sum = (sum + arr[i]) % K

		// If sum is negative, make it positive by adding K
		/**
		Verifying the Positive Remainder: Now, to verify that 2 is indeed the correct positive
		 equivalent of -1 for the expression -4 % 3:
		Consider the number -4 and add 2 to it, yielding -2. While -2 itself is not a multiple
		of 3, the relationship we are interested in is the difference between -4 and the nearest
		higher multiple of 3.
		The nearest higher multiple of 3 to -4 is 0. The difference between -4 and 0 is 4, which
		is 1 more than 3. This extra 1 is effectively the positive counterpart of the original remainder -1.
		Why Adding K Works: The addition of K to the negative remainder effectively "rolls" the number
		 up to the next multiple of K. In our case, it rolls -4 up to the range of the next multiple of 3,
		 which is 0. The new remainder, 2, is the positive version of -1, and it correctly represents the
		 remainder of dividing -4 by 3 in a non-negative form.
		*/
		if sum < 0 {
			sum = sum + K
		}
		if modIdx, ok := prefixMod[sum]; ok {
			length := i - modIdx
			if length > longestLength {
				longestLength = length
				count = 1 // reset count for a new longest length
			} else if length == longestLength {
				count++
			}
		} else {
			prefixMod[sum] = i
		}
	}

	return count
}

func main() {

	// input := `4 3
	// 2 3 4 6`
	// input := `3 3
	// 3 -4 5`

	input := `5 3
2 4 3 5 1`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	// scanner.Scan()
	// t, _ := strconv.Atoi(scanner.Text())

	// for i := 0; i < t; i++ {
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
	fmt.Println(countSubarrayAndSum(arr, N, K))
	// }
}
