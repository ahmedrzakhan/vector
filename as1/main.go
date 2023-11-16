package main

import (
	"fmt"
	"strconv"
)

// isPalindrome checks if a string is a palindrome.
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Println("Enter an integer:")
	fmt.Scan(&num)

	// Convert integer to string
	strNum := strconv.Itoa(num)

	// Check if the string representation of the number is a palindrome
	if isPalindrome(strNum) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
