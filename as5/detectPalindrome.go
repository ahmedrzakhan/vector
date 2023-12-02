package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
given a string detect if it can be converted into palindrome with any rearrangeement
deletion addition not allowed
*/

// TC - O(N) SC- O(N)
func canBePalindrom(str string, N int) bool {
	strMap := make(map[rune]int)

	for _, ch := range str {
		strMap[ch]++
	}

	oddCount := 0
	for _, count := range strMap {
		if count%2 == 1 {
			oddCount++
		}
		if oddCount > 1 {
			return false
		}
	}
	return true
}

func mainDetectPalindrome() {
	input := `2
	6
	aabbcc
	5
	aabcd`
	// Possible!
	// Not Possible!
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < T; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		str := scanner.Text()
		pali := canBePalindrom(str, N)
		if pali {
			fmt.Println("Possible!")
		} else {
			fmt.Println("Not Possible!")
		}
	}
}
