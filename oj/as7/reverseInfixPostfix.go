package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
if letter push it into stack, if operator pop last 2 items in stack
pad them with () and keep operatot in middle
*/

func isOperator(token rune) bool {
	return strings.ContainsRune("+-*/^", token)
}

// TC - O(N), SC - O(N)
func postfixToInfix(postfix string) string {
	stack := []string{}
	for _, token := range postfix {
		fmt.Println("token", string(token))
		if unicode.IsLetter(token) {
			stack = append(stack, string(token))
		} else if isOperator(token) {
			op2 := stack[len(stack)-1]   // Top of stack
			stack = stack[:len(stack)-1] // Pop
			op1 := stack[len(stack)-1]   // New top of stack
			stack = stack[:len(stack)-1] // Pop
			// Combine and push back to stack
			stack = append(stack, fmt.Sprintf("(%s%c%s)", op1, token, op2))
		}
	}

	if len(stack) > 0 {
		return stack[len(stack)-1]
	}

	return ""
}

func mainRev() {
	postfix := "abc*+d-"
	// ((a+(b*c))-d)
	// scanner := bufio.NewScanner(os.Stdin)
	// ((((a+b)-c)+((d*(e-f))/g))+(h*(i/j)))
	// scanner.Scan()
	// postfix := scanner.Text()
	// postfix := "ab+c-def-*g/+hij/*+"
	// ((((a+b)-c)+((d*(e-f))/g))+(h*(i/j)))
	fmt.Println(postfixToInfix(postfix))
}
