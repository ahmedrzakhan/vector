package main

import (
	"fmt"
	"strconv"
)

/**
if palindrome  print yes else no
input 1221 output yes
*/

func isPalindrome(s string) bool {
	lengthOfStr := len(s)

	for i := 0; i < lengthOfStr/2; i++ {
		if s[i] != s[lengthOfStr-1-i] {
			return false
		}
	}

	return true
}

func mainisPalindrome() {
	num := 1221

	// Convert integer to string
	strNum := strconv.Itoa(num)

	// Check if the string representation of the number is a palindrome
	if isPalindrome(strNum) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
