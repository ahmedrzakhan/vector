package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Given an array, find the nearest smaller element G[i] for every element A[i] in the array such that the element
has an index smaller than i. Mathematically, G[i] for an element A[i] is an element A[j] such that j is maximum
possible AND j < i AND A[j] < A[i]
Note: Elements for which no smaller element exist, consider next smaller element as -1.
*/

// TC O(N^2) SC O(N)
func getNearest(arr []int) string {
	N := len(arr)
	var response []int
	for i := N - 1; i >= 0; i-- {
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] < arr[i] {
				response = append(response, arr[j])
				break
			}
		}
		if j == -1 {
			response = append(response, -1)
		}
	}

	return strings.Trim(fmt.Sprint(reverse(response)), "[]")
}

func reverse(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func smallerNeighbourElement(arr []int) string {
	// Create a stack and a result slice with the same length as the input array.
	stack := make([]int, 0)
	result := make([]int, len(arr))

	// Initialize result array with -1 as default value.
	for i := range result {
		result[i] = -1
	}

	for i, value := range arr {
		// Pop elements from the stack until the current element is greater
		// than the element at the top of the stack.
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= value {
			stack = stack[:len(stack)-1]
		}
		// If stack is not empty, the top element is the nearest smaller element.
		if len(stack) > 0 {
			result[i] = arr[stack[len(stack)-1]]
		}
		// Push the current index onto the stack.
		stack = append(stack, i)
	}

	return strings.Trim((fmt.Sprint(result)), "[]")
}

func mainSmaller() {
	input := `8
39 27 11 4 24 32 32 1`
	// -1 -1 -1 -1 4 24 24 -1

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

	fmt.Println(smallerNeighbourElement(arr))
}
