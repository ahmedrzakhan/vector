package main

import (
	"bufio"
	"fmt"
	"strings"
)

/**
Strings with Parenthesis

Description:
Given a string containing three types of parentheses: (), [], {}, write a program to check whether the string contains a valid sequence of parentheses.

Input:
Input Format: Input contains one line which is the string with parentheses whose validity you have to check.
Constraints: length of string <= 1000.

Output:
Output Format: Print "balanced" (without quotes) if it has a sequence of valid parentheses, otherwise print "unbalanced" (without quotes).

Sample Input 1:

bash
Copy code
(((((((((())))))))))
Sample Output 1:

Copy code
balanced
Sample Input 2:

csharp
Copy code
[one[two[three[four[five[six[seven[eight[nine[ten]]]]]]]]]]
Sample Output 2:

Copy code
balanced
*/

/**
If number of items odd in total return unbalanced immediately at the start

Ask interviewer if anything other than these six characters
*/

func includes(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// TC - O(N) SC - O(N)
func getStatus(s string) bool {
	elemsToInsert := []string{"(", "{", "["}
	elemsToRemove := []string{")", "}", "]"}

	var stack []string

	for i := 0; i < len(s); i++ {
		ele := string(s[i])
		if includes(elemsToInsert, ele) {
			stack = append(stack, string(s[i]))
		} else if includes(elemsToRemove, ele) {
			if len(stack) == 0 {
				return false
			} else {
				lastELem := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if ele == ")" && lastELem != "(" || ele == "}" && lastELem != "{" ||
					ele == "]" && lastELem != "[" {
					return false
				}
			}
		}

	}

	return len(stack) == 0
}

func mainPa() {
	// input := `{(([])[])[]]}`
	input := `(((((((((())))))))))`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	status := getStatus(s)
	if status {
		fmt.Println("balanced")
	} else {
		fmt.Println("unbalanced")
	}

}
