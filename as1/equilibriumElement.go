package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/*
Given an array A of N positive numbers. The task is to find the position where
equilibrium first occurs in the array. An equilibrium position in an array is
a position such that the sum of elements before it is equal to the sum of elements
after it. The valid index range is from [ 1, n-2 ] because there should be at least
one element on both sides.
*/
func findEquilibrium(arr []int) int {
	totalSum := 0
	for _, num := range arr {
		totalSum += num
	}

	leftSum := 0
	for i, num := range arr {
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}
	return -1
}

func mainEq() {
	input := `2
	5
	15 1 5 5 5
	3
	1 2 3`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text()) // Number of test cases

	for i := 0; i < T; i++ {
		scanner.Scan()
		// N, _ := strconv.Atoi(scanner.Text()) // Size of the array, not used here
		scanner.Scan()
		nums := strings.Fields(scanner.Text())
		arr := make([]int, len(nums))

		for j, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			arr[j] = num
		}

		fmt.Println(findEquilibrium(arr))
	}
}
