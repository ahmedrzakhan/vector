package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
For each test case, print the number of elements common in both the arrays on a single line,
in ascending order -1, if no elements are common between the two arrays, on a new line
*/

// TC - O(N+M)
func getCommon(arrA, arrB []int, NA, NB int) []int {
	indexA, indexB := 0, 0
	var commonElems []int
	for indexA < len(arrA) && indexB < len(arrB) {
		if arrA[indexA] == arrB[indexB] {
			commonElems = append(commonElems, arrA[indexA])
			indexA++
			indexB++
		} else if arrA[indexA] < arrB[indexB] {
			indexA++
		} else {
			indexB++
		}
	}

	if len(commonElems) == 0 {
		return []int{-1}
	}
	return commonElems
}

func mainCom() {
	// 	input := `2
	// 6
	// 1 2 3 4 5 6
	// 3
	// 3 3 5
	// 4
	// 1 2 3 4
	// 4
	// 5 6 7 8`

	// 	reader := strings.NewReader(input)
	// 	scanner := bufio.NewScanner(reader)
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < T; i++ {
		scanner.Scan()
		NA, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		arrStrA := strings.Fields(scanner.Text())

		arrA := make([]int, NA)
		for index, elem := range arrStrA {
			arrA[index], _ = strconv.Atoi(elem)
		}

		scanner.Scan()
		NB, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		arrStrB := strings.Fields(scanner.Text())
		arrB := make([]int, NB)

		for index, elem := range arrStrB {
			arrB[index], _ = strconv.Atoi(elem)
		}

		common := getCommon(arrA, arrB, NA, NB)
		fmt.Println(strings.Trim(fmt.Sprint(common), "[]"))
	}
}
