package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Given a binary array A of size n, find the maximum length of a contiguous subarray with
equal number of 0 and 1.
*/

// TC -O(N^2) SC -O(1)
func getCountBr(arr []int, N int) int {
	maxLength := 0
	for i := 0; i < N; i++ {
		countZero, countOne := 0, 0
		j := i
		for ; j < N; j++ {
			if arr[j] == 0 {
				countZero++
			} else {
				countOne++
			}
			if countZero == countOne {
				maxLength = max(maxLength, j-i+1)
			}
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TC O(N) SC- O(1)
func getCountOpt(arr []int, N int) int {
	maxLength, sum := 0, 0
	sumIndexMap := map[int]int{0: -1}

	for i := 0; i < N; i++ {
		if arr[i] == 0 {
			sum--
		} else {
			sum++
		}

		if prevIndex, exists := sumIndexMap[sum]; exists {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			sumIndexMap[sum] = i
		}
	}
	return maxLength
}

func mainSub() {
	input := `1
	5
	1 0 0 1 0`

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

		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())
		arr := make([]int, N)
		for index, elem := range arrStr {
			arr[index], _ = strconv.Atoi(elem)
		}

		count := getCountOpt(arr, N)
		fmt.Println(count)
	}

}
