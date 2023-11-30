package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLongestCount(s []int, N int) int {
	count, maxCount := 0, 0

	for i := 0; i < N; i++ {
		if s[i]%2 == 1 {
			count++
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}

	return maxCount
}

func mainLongest() {
	// reader := strings.NewReader(input)
	// scanner := bufio.NewScanner(reader)
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())
	arr := make([]int, N)

	for index, elem := range arrStr {
		arr[index], _ = strconv.Atoi(elem)
	}

	count := getLongestCount(arr, 12)
	fmt.Println(count)
}
