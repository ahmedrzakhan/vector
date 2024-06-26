// Enter code here
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Minions are of 26 types, from 'a' to 'z'. If two Minions of same type finds themselves besides each
other, they will kill each other.</p><p>You are given a string of Minions. You need to find the final
 surviving string of Minions, after all the killings. i.e Further No Minion will kill other Minion.
 If no Minion survived a after all the fighting, print "-1".
*/

// TC -O(N) SC - O(N)
func getMinionss(s string) string {
	var stack []string

	for i := 0; i < len(s); i++ {
		if len(stack) != 0 && stack[len(stack)-1] == string(s[i]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(s[i]))
		}
	}

	return strings.Join(stack, "")
}

func getMinions(s string) string {
	var stack []rune
	// NOTE: range gets rune, while basic for gets string
	for _, ch := range s {
		if len(stack) != 0 && stack[len(stack)-1] == ch {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return string(stack)
}

func mainW() {
	// 	input := `5
	// aabbc`

	// 	reader := strings.NewReader(input)
	// 	scanner := bufio.NewScanner(reader)
	scanner := bufio.NewScanner(os.Stdin)
	// NOTE: bigger test cases were giving wrong answer, coz of buffer size
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	_, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	str := scanner.Text()

	resp := getMinions(str)
	if len(resp) == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(resp)
	}
}
