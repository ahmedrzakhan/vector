package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
For the first day, span is always 1. In the example we can see that for day 2 at 60,
there is no day before it where the price was less than 60. Hence span is 1 again.
For day 3, the price at day 2 (60) is less than 70, hence span is 2. Similarly, for
day 4 and day 5. Remember days should be consecutive, that’s why the span for day 4
is 1 – even though there was a day 2 where the price was less than 65.
*/

// TC - O(N), SC - O(N)
func getFinance(arr []int) string {
	stack := []int{}
	result := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		for len(stack) != 0 && arr[stack[len(stack)-1]] < arr[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			result[i] = i + 1
		} else {
			result[i] = i - stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return strings.Trim(fmt.Sprint(result), "[]")
}

func mainFina() {
	input := `1
6
100 60 70 65 80 85`
	// 1 1 2 1 4 5

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi((scanner.Text()))

	for i := 0; i < T; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())
		arr := make([]int, N)

		for i, ele := range arrStr {
			arr[i], _ = strconv.Atoi(ele)
		}
		fmt.Println(getFinance(arr))
	}
}
