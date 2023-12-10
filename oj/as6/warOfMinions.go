// Enter code here
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TC -O(N) SC - O(N)
func getMinions(s string) string {
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

func mainW() {
	// 	input := `5
	// aabbc`

	// 	reader := strings.NewReader(input)
	// 	scanner := bufio.NewScanner(reader)
	scanner := bufio.NewScanner(os.Stdin)
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
