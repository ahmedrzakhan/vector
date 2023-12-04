package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Given an array A, print starting and ending index of all subarrays in the array which has sum 0. (0-indexing)
*/

// TC - O(N^2) SC - O(N)
func getZeroSum(arr []int) {
	N := len(arr)
	for i := 0; i < N; i++ {
		if arr[i]+arr[i] == 0 {
			fmt.Println(i, i)
		}
		sum := arr[i]
		for j := i + 1; j < N; j++ {
			sum = sum + arr[j]
			if sum == 0 {
				fmt.Println(i, j)
			}
		}
		sum = 0
	}

}

func mainSubarray() {
	// arr := []int{6, 3, -1, -3, 4, -2, 2, 4, 6}
	input := `6 3 -1 -3 4 -2 2 4 6`
	// 2 4
	// 2 6
	// 5 6
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())

	var arr []int

	for _, ele := range arrStr {
		num, _ := strconv.Atoi(ele)
		arr = append(arr, num)
	}

	getZeroSum(arr)
}
