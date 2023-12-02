package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Given a array A having N positive integers. Count all the subarrays of A
having length X, such that each subarray has no integer greater than K
*/

// TC - O(N^2) SC - O(1)
func getCountOfSubarraysB(arr []int, N int, K int, X int) int {
	fmt.Println("arr", arr)
	count := 0
	for i := 0; i < N-X+1; i++ {
		invalid := false
		for j := i; j < i+X; j++ {
			if arr[j] > K {
				invalid = true
				break
			}
		}
		if !invalid {
			count++
		}
	}
	return count
}

// TC - O(N) SC - O(1)
func getCountOfSubarrays(arr []int, N int, K int, X int) int {
	count := 0
	exceed := 0 // Number of elements greater than K in the current window

	// Initialize the count of exceed for the first window
	for i := 0; i < X; i++ {
		if arr[i] > K {
			exceed++
		}
	}

	// If the first window doesn't exceed, count it
	if exceed == 0 {
		count++
	}

	// Slide the window and update counts
	for i := X; i < N; i++ {
		// Element exiting the window
		if arr[i-X] > K {
			exceed--
		}

		// Element entering the window
		if arr[i] > K {
			exceed++
		}

		// If current window doesn't exceed, count it
		if exceed == 0 {
			count++
		}
	}

	return count
}

func mainXSubarray() {
	input := `1
	6 3 2
	2 1 3 4 2 1`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < T; i++ {
		scanner.Scan()
		NKX := strings.Fields(scanner.Text())
		N, _ := strconv.Atoi(NKX[0])
		K, _ := strconv.Atoi(NKX[1])
		X, _ := strconv.Atoi(NKX[2])

		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())
		arr := make([]int, N)
		for i, ele := range arrStr {
			arr[i], _ = strconv.Atoi(ele)
		}

		count := getCountOfSubarrays(arr, N, K, X)
		fmt.Println(count)
	}
}
