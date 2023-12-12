package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Signal's Capacity

Description:
There are many signal towers present in Bangalore. Towers are aligned in a straight horizontal line (from left to right) and each tower transmits a signal in the right to left direction.

Tower X shall block the signal of Tower Y if Tower X is present to the left of Tower Y and Tower X is taller than Tower Y. So, the power of a signal of a given tower can be defined as: {(the number of contiguous towers just to the left of the given tower whose height is less than or equal to the height of the given tower) + 1}.

You need to write a program that finds the power of each tower.

Input:
Input Format:
Each test case has multiple test cases in it:

First line contains an integer specifying the number of test cases.
Second line contains an integer n specifying the number of towers.
Third line contains n space-separated integers (H[i]) denoting the height of each tower.
Constraints:

1 <= T <= 10
2 <= n <= 10^6
1 <= H[i] <= 10^8
Output:
Output Format:
Print the range of each tower (separated by a space).

Sample Input 1:
2
7
100 80 60 70 60 75 85
5
3 5 0 9 8

Sample Output 1:
1 1 1 2 1 4 6
1 2 1 4 1

Hint:
Sample 1 Explanation:
There are 2 test cases:
In the first test case:
7 towers are present, and the signal range of each tower is separated by space: 1 1 1 2 1 4 6.

Similarly, in the second test case, we have 5 towers whose signal range is given.
*/

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
