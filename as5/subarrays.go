package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
You are given an array A of N elements. Write a program that counts the number of sub-arrays of A in which sum of all the elements is an even number
*/

// TC - O(N^2) SC - O(1)
func getSumM(arr []int, N int) int {
	count := 0
	for i := 0; i < N; i++ {
		sum := 0
		for j := i; j < N; j++ {
			sum = sum + arr[j]
			if sum%2 == 0 {
				count++
			}
		}
	}
	return count
}

func mainSum() {
	input := `5
	2 5 4 4 4`
	// 7
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
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
	sum := getSumM(arr, N)
	fmt.Println(sum)
}
