package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// TC - O(N) SC O(1)
func getLongestIncreasingSubarrayCount(s []int, N int) int {
	count, maxCount, prevOdd := 0, 0, -1

	for i := 0; i < N; i++ {
		if s[i]%2 == 1 {
			if s[i] == prevOdd {
				count++
			} else {
				prevOdd = s[i]
				count = 1
			}
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
			prevOdd = -1
		}
	}

	return maxCount
}

func mainLongestInc() {
	input := `2
2
1 1
6
1 2 1 2 3 1`
	// output
	// 1
	// 3

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < T; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())
		arr := make([]int, N)
		for index, elem := range arrStr {
			arr[index], _ = strconv.Atoi(elem)
		}

		count := getLongestIncreasingSubarrayCount(arr, N)
		fmt.Println(count)
	}

}
