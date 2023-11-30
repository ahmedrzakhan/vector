package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
two arrays one sorted in ascending other in descending, tell number of common elements
in arrays without extra space
*/

// TC - O(N)
func getCount(A []int, B []int, N int) int {
	L, R, count := 0, N-1, 0
	for L <= N-1 && R >= 0 {
		if A[L] == B[R] {
			count++
			L++
			R--
		} else if A[L] < B[R] {
			L++
		} else if A[L] > B[R] {
			R--
		}
	}

	return count
}

func mainTwosorted() {

	// a := []int{1, 2, 2, 3, 4, 5}
	// b := []int{4, 4, 3, 2, 1, 1}

	input := `1
6
1 2 2 3 4 5
4 4 3 2 1 1`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()

		arrAStr := strings.Fields(scanner.Text())
		scanner.Scan()
		arrBStr := strings.Fields(scanner.Text())
		arrA := make([]int, N)
		arrB := make([]int, N)
		for index, elem := range arrAStr {
			arrA[index], _ = strconv.Atoi(elem)
		}
		for index, elem := range arrBStr {
			arrB[index], _ = strconv.Atoi(elem)
		}
		count := getCount(arrA, arrB, N)
		fmt.Println("print", count)
	}

}
