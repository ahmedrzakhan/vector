package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
You are given an array of N integers and a single integer K. You need to find out if there is a subarray, which has the sum exactly as K
*/

func getCountA(arr []int, N, K int) string {
	sum := 0
	sumMap := make(map[int]int)

	// Initialize the sumMap for sum 0 at index -1 to handle cases where the subarray starts at index 0
	sumMap[0] = -1

	for i := 0; i < N; i++ {
		sum += arr[i]

		// If the current sum equals K, we've found a subarray from index 0 to i
		if sum == K {
			return "Yes"
		}

		// Check if there's a prefix sum that, when removed from the current sum, results in K
		if _, ok := sumMap[sum-K]; ok {
			return "Yes"
		}

		// Only set the index in the map if this sum hasn't been seen before to ensure we always have the earliest index
		if _, exists := sumMap[sum]; !exists {
			sumMap[sum] = i
		}
	}

	return "No"
}

func mainSumOf() {
	input := `1
	3 1
	1 3 2`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)      // create a buffer with a 64KB initial size
	scanner.Buffer(buf, 10*1024*1024)    // allow the buffer to grow to a maximum of 10MB
	scanner.Scan()                       // Read the first line to get the number of test cases
	t, _ := strconv.Atoi(scanner.Text()) // Convert the number of test cases to an integer

	for i := 0; i < t; i++ {
		scanner.Scan()                       // Read the line containing N and K
		NK := strings.Fields(scanner.Text()) // Split the line into fields
		N, _ := strconv.Atoi(NK[0])          // Convert N to an integer
		K, _ := strconv.Atoi(NK[1])          // Convert K to an integer

		scanner.Scan()                           // Read the line containing the array elements
		arrStr := strings.Fields(scanner.Text()) // Split the line into fields
		arr := make([]int, N)
		for index, elem := range arrStr {
			arr[index], _ = strconv.Atoi(elem) // Convert each array element to an integer
		}
		res := getCountA(arr, N, K)
		fmt.Println(res)
	}
}
