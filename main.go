package main

import "fmt"

/**
Given an array of integers temperatures represents the daily temperatures, return an array answer
such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.
Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]
*/

func main() {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	// arr := []int{30, 60, 90}

	N := len(arr)
	ans := make([]int, N)
	var stack []int // stack to keep indices of temperatures

	for i, temp := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] < temp {
			// The current day's temperature is warmer than the temperature
			// at the index on the top of the stack.
			// So we can pop from the stack and calculate the difference in days
			// between the current index and the index from the stack.
			lastElem := stack[len(stack)-1] // Get the last element from the stack
			stack = stack[:len(stack)-1]    // Pop from the stack
			ans[lastElem] = i - lastElem    // Calculate the difference
		}
		// Push the current index onto the stack
		stack = append(stack, i)
	}

	// Remaining indices in the stack don't have warmer temperatures in the future
	// The 'ans' slice is already initialized with 0s, so we don't have to explicitly set them.
	fmt.Println(ans)
}
