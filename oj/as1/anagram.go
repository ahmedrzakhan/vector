package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
Given an integer n, we have n strings of length 4. Now we have to remove all those strings which is an anagram of previously appeared string
*/
// Helper function to sort a string
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	input := `5
	eodc
	odec
	code
	baca
	aabc`
	// baca
	// eodc
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	seen := make(map[string]bool)
	var result []string

	for i := 0; i < n; i++ {
		scanner.Scan()
		str := scanner.Text()
		sortedStr := sortString(str)

		if !seen[sortedStr] {
			seen[sortedStr] = true
			result = append(result, str)
		}
	}

	sort.Strings(result)
	fmt.Println(len(result))
	for _, str := range result {
		fmt.Println(str)
	}
}
