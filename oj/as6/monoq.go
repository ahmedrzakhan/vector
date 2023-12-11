package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
You are given a string of length N. You need to find out the lexicographic minimum subsequence of length K from that
string. A subsequence is a resulting sequence obtained by erasing some (possibly zero) characters from the string.
*/

// TC - O(N) SC - O(N)
func getString(str string, K int) string {
	var stack []rune

	// Iterate over each character
	for i, ch := range str {
		// While the stack is not empty and the top of the stack is greater than
		// the current character, and we have enough characters remaining to reach
		// a length of 'k', pop the stack
		for len(stack) > 0 && stack[len(stack)-1] > ch && len(str)-i+len(stack) > K {
			stack = stack[:len(stack)-1]
		}
		// If the stack length is less than 'k', push the current character
		if len(stack) < K {
			stack = append(stack, ch)
		}
	}

	return string(stack)
}

func mainSt() {
	input := `12 6
eyaaaaaaaaaa`
	// egztkh

	// 	4 2
	// dbac
	// ac

	// 3 3
	// abc

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// NOTE: bigger test cases were giving wrong answer, coz of buffer size
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	NK := strings.Fields(scanner.Text())
	_, _ = strconv.Atoi(NK[0])
	K, _ := strconv.Atoi(NK[1])

	scanner.Scan()
	str := scanner.Text()

	fmt.Println(getString(str, K))
}
