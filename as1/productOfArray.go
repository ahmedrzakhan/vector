package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	// Initialize the answer array
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	// fmt.Println("answer", answer)
	// fmt.Println("nums", nums)

	// Multiply from the right
	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		// fmt.Println("R before", R)
		R *= nums[i]
		// fmt.Println("num[i]", nums[i])
		// fmt.Println("R after", R)
	}
	// fmt.Println("answer", answer)
	// fmt.Println("nums", nums)

	return answer
}

func mainProductExceptSelf() {

	nums := []int{2, 3, 4, 5}

	result := productExceptSelf(nums)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
