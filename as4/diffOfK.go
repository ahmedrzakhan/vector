package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/*
* diff of two should equal to K
* array in ascending sorted order
* also contains negative integers
 */
func getDiff(arr []int, N, K int) string {
	left, right := 0, 1

	for right < len(arr) {
		diff := arr[right] - arr[left]
		if diff == K {
			return "Yes"
		} else if diff < K {
			right++
		} else {
			left++
			if left == right {
				right++
			}
		}
	}
	return "No"
}

func mainDiff() {
	input := `2
	6 0
	-8 -7 5 6 6 9
	5 8
	1 2 3 4 5`
	// Yes
	// No

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

		resp := getDiff(arr, N, K)
		fmt.Println(resp)
	}
}
