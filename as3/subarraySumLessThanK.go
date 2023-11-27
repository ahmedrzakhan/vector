package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// TC - O(N^2) SC - O(1)
func countOfSubarraysLessThanKBrute(arr []int, N int, K int) int {
	count := 0
	for i := 0; i < N; i++ {
		sum := 0
		for j := i; j < N; j++ {
			sum = sum + arr[j]
			if sum < K {
				count++
			} else {
				break
			}
		}

	}

	return count
}

// Optimized TC - O(N), SC - O(1)
func countOfSubarraysLessThanK(arr []int, N int, K int) int {
	count := 0
	start := 0
	sum := 0

	for end := 0; end < N; end++ {
		// Add the current element to the sum
		sum += arr[end]

		// While the sum is at least K, remove the starting elements
		for sum >= K {
			sum -= arr[start]
			start++
		}

		// All subarrays between start and end have a sum less than K
		count = count + end - start + 1
	}

	return count
}

func mainCountOfSubarraysLessThanKBrute() {

	input := `1
	5 5
	1 5 1 3 2`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
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
		fmt.Println(countOfSubarraysLessThanK(arr, N, K))
	}
}
