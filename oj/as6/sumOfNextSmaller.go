package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Description

Given an array of integers, find the sum of the nearest smaller values to the right of all the elements, if such values don't exist, consider them zero.

Input

First-line contains T, no of test cases.
First-line of each test case contains n which is the number of elements in the array
Second-line of each test case contains n numbers representing elements of the array

Constraints

1 <= T <= 10
1 <= N <= 100000
1 <= Ai <= 100

Output

For each test case, output a single integer, the answer to the problem, on a new line

Sample Input 1

2
4
4 3 1 2
4
1 2 3 4

Sample Output 1

4
0

Hint

For 1st test case:-
For 4 first smaller value to the right is 3.
For 3 first smaller value to the right is 1.
For 1 first smaller value to the right is not there.
For 2 first smaller value to the right is not there.

6 5 6 2
in this case 2 gets added twice
which was acceptable in test cases
*/

// TC O(N) SC O(N)
func getSum(arr []int) int {
	sum := 0
	stack := make([]int, 0)
	for i, num := range arr {
		for len(stack) != 0 && arr[stack[len(stack)-1]] > num {
			stack = stack[:len(stack)-1]
			sum = sum + num
		}
		stack = append(stack, i)
	}
	return sum
}

// 4 3 1 2

func mainsm() {
	input := `3
4
6 5 6 2
4
1 2 3 4
8
6 5 1 4 2 1 3 2`

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

		fmt.Println(getSum(arr))
	}
}
