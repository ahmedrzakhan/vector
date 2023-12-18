package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Minions are of 26 types, from 'a' to 'z'. If two Minions of same type finds themselves besides each other, they will kill each other.
You are given a string of Minions. You need to find the final surviving string of Minions, after all the killings. i.e Further No
Minion will kill other Minion.If no Minion survived a after all the fighting, print "-1".
5
aabbc
-> c
12
abbabaadcbbc
-> bd
2
aa
-> -1
*/

func getMinions(str string) string {
	N := len(str)
	stack := []string{}
	for i := 0; i < N; i++ {

		if len(stack) != 0 && stack[len(stack)-1] == string(str[i]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(str[i]))
		}

	}
	if len(stack) == 0 {
		return "-1"
	}

	return strings.Join(stack, "")
}

func mainMinion() {
	input := `12
abbabaadcbbc`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	_, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()

	str := scanner.Text()
	fmt.Println(getMinions(str))
}
