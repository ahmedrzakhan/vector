package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/*
*
You have to find the minimum number of adjacent swaps required to make the string palindrome.
*/
func canBePalindrome(runes []rune) bool {
	counts := make(map[rune]int)
	oddCount := 0
	for _, r := range runes {
		counts[r]++
	}

	for _, count := range counts {
		if count%2 == 1 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}
	return true
}

func swap(runes []rune, i, j int) {
	runes[i], runes[j] = runes[j], runes[i]
}

// TC O(N^2)
func minSwapsToPalindrome(s string) int {
	runes := []rune(s)
	swapCount := 0

	if !canBePalindrome(runes) {
		return -1
	}

	for i := 0; i < len(runes)-1; i++ {
		j := len(runes) - 1 - i
		for ; runes[i] != runes[j] && j > i; j-- {
		}
		for k := j; k < len(runes)-1-i; k++ {
			swap(runes, k, k+1)
			swapCount++
		}
	}

	return swapCount
}

func mainSwaps() {
	input := `2
5
aabcb
8
adbcdbad`
	// 3
	// -1

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 64*1024*1024)
	scanner.Scan()

	T, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < T; i++ {
		scanner.Scan()
		_, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		fmt.Println(minSwapsToPalindrome(scanner.Text()))

	}

}
