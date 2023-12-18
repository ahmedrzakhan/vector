package main

import (
	"bufio"
	"fmt"
	"strings"
)

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	default:
		return 0
	}
}

// TC - O(N), SC - O(N)
func infixToPostfix(infix string) string {
	stack := []rune{}
	var postfix strings.Builder

	for _, token := range infix {
		switch {
		case token >= 'a' && token <= 'z':
			postfix.WriteRune(token)
		case token == '(':
			stack = append(stack, token)
		case token == ')':
			for len(stack) != 0 && stack[len(stack)-1] != '(' {
				postfix.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		default:
			for len(stack) != 0 && precedence(stack[len(stack)-1]) >= precedence(token) {
				postfix.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		}
	}

	for len(stack) != 0 {
		postfix.WriteRune(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix.String()
}

func mainInfix() {
	// infix
	input := `a+b*(c^d-e)^(f+g*h)-i`
	// ab+c-def-*g/+hij/*+
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	infix := scanner.Text()

	fmt.Println(infixToPostfix(infix))
}
