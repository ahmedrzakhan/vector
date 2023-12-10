package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// findMinXOR finds the minimum XOR value between any two elements in the array.
func findMinXOR(arr []int) int {
	// Sort the array
	sort.Ints(arr)

	// Initialize min XOR value with the maximum possible value.
	minXOR := int(^uint(0) >> 1) // This sets minXOR to the maximum positive int value.

	// Compute the XOR of each adjacent pair and find the minimum.
	for i := 0; i < len(arr)-1; i++ {
		xor := arr[i] ^ arr[i+1]
		if xor < minXOR {
			minXOR = xor
		}
	}

	return minXOR
}

func mainXor() {
	input := `1
4
0 2 5 7`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < T; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())

		arr := make([]int, N)
		for index, elem := range arrStr {
			arr[index], _ = strconv.Atoi(elem)
		}
		fmt.Println(findMinXOR(arr))
	}
}
