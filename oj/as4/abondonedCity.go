package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
given array fin minimum subarray need to get sum K
The output should contain the shortest distance you need to walk to collect target coins
*/

func getSmallestSubarrayCount(arr []int, N int, K int) int {
	start, end, sum, minLen := 0, 0, 0, N+1

	for end < N {
		// Expand the window by including more elements to the sum
		sum = sum + arr[end]

		// Shrink the window from the start as long as the sum is greater than or equal to K
		for sum >= K {
			if minLen > (end - start + 1) {
				minLen = end - start + 1
			}
			sum -= arr[start]
			start++
		}

		end++
	}

	// If minLen is not updated then no subarray is found
	if minLen == N+1 {
		return -1
	}

	return minLen
}

func mainAbon() {
	input := `1
	5 7
	1 2 3 4 5`
	// 	input := `2
	// 5 13
	// 5 1 2 3 4
	// 1 10
	// 1`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()

	T, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < T; i++ {
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
		count := getSmallestSubarrayCount(arr, N, K)
		fmt.Println(count)
	}

}
