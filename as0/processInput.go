package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainProcessInput() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)      // create a buffer with a 64KB initial size
	scanner.Buffer(buf, 10*1024*1024)    // allow the buffer to grow to a maximum of 10MB
	scanner.Scan()                       // Read the first line to get the number of test cases
	t, _ := strconv.Atoi(scanner.Text()) // Convert the number of test cases to an integer

	for i := 0; i < t; i++ {
		scanner.Scan()                       // Read the line containing N and K
		nk := strings.Fields(scanner.Text()) // Split the line into fields
		N, _ := strconv.Atoi(nk[0])          // Convert N to an integer
		K, _ := strconv.Atoi(nk[1])          // Convert K to an integer

		scanner.Scan()                           // Read the line containing the array elements
		arrStr := strings.Fields(scanner.Text()) // Split the line into fields
		arr := make([]int, N)
		for j, s := range arrStr {
			arr[j], _ = strconv.Atoi(s) // Convert each array element to an integer
		}

		// Now, arr is the array and k is the sum we're looking for
		// You can now call a function to check for a subarray with sum k
		// For example: hasSubarrayWithSum(arr, k)

		fmt.Println("Read array:", arr, "with sum:", K)
	}
}

// Define the function hasSubarrayWithSum based on the earlier provided algorithm

/**
3
5 3
1 2 1 3 4
4 5
1 2 1 3
3 2
1 2 1
*/
