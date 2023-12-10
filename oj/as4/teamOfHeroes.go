package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/*
* sum of two should equal to K
* array in ascendong sorted order
 */
func getHero(arr []int, N, K int) string {
	left, right := 0, N-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == K {
			return "Yes"
		} else if sum < K {
			left++
		} else {
			right--
		}
	}

	return "No"
}

func mainHroe() {
	input := `2
	5 31
	10 11 13 17 21
	5 44
	10 11 13 17 21`
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

		resp := getHero(arr, N, K)
		fmt.Println(resp)
	}
}
