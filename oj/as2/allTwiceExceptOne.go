package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Given an array A of n elements, every element appears twice in the array except for one
element. Find the element that is uniquely present in the array
*/

/**
NOTE:
This is a classic problem that can be solved using the XOR bitwise operation. When you XOR
a number with itself, the result is 0. And when you XOR any number with 0, the result is
 the number itself. So, by XORing all the numbers in the array together, the result will
 be the unique number that is not repeated.
*/

func findUniqueElement(nums []int) int {
	unique := 0
	for _, num := range nums {
		unique ^= num
	}
	return unique
}

func mainAll() {
	input := `2
1
5
5
1 2 1 3 2`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < T; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		arr := make([]int, N)
		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())
		for j, elem := range arrStr {
			arr[j], _ = strconv.Atoi(elem)
		}
		fmt.Println(findUniqueElement(arr))
	}
}
