package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// TC - O(N) SC - O(N)
func calculateSignalPower(arr []int) string {
	stack := []int{} // Stack to hold indices of towers
	result := make([]int, len(arr))

	for i, ele := range arr {
		// Start with power of 1 for the current tower
		sum := 1
		// Pop from stack until a taller tower is found
		for len(stack) > 0 && arr[stack[len(stack)-1]] <= ele {
			sum = sum + result[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
		}
		result[i] = sum
		stack = append(stack, i) // Push the index of the current tower
	}

	return strings.Trim(fmt.Sprint(result), "[]")
}

func mainSignal() {
	input := `2
7
100 80 60 70 60 75 85
5
3 5 0 9 8`

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
		for i, ele := range arrStr {
			arr[i], _ = strconv.Atoi(ele)
		}

		fmt.Println(calculateSignalPower(arr))
	}
}
