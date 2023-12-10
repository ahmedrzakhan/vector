package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
2 arrays diff size
if both have same elems return yes else no
*/

func removeDuplicate(arr []int, N int) []int {
	var newArr []int
	for i := 0; i < N; i++ {
		if i != N-1 && arr[i] != arr[i+1] {
			newArr = append(newArr, arr[i])
		}
	}

	return newArr
}

// TC O(NlogN + MlogM) SC O(N + M)
func isSame(arr, brr []int, AN, BN int) string {
	sort.Ints(arr)
	sort.Ints(brr)

	newArr := removeDuplicate(arr, AN)
	newBrr := removeDuplicate(brr, BN)

	if len(newArr) != len(newBrr) {
		return "No"
	}

	for i := 0; i < len(newArr); i++ {
		if newArr[i] != newBrr[i] {
			return "No"
		}
	}

	return "Yes"
}

func mainParty() {
	//	input := `6
	//
	// 1 2 2 3 4 5
	// 6
	// 1 2 3 4 5 5`
	input := `5
1 2 2 3 4
6
1 2 3 4 5 5`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	AN, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())

	scanner.Scan()
	BN, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	brrStr := strings.Fields(scanner.Text())

	arr := make([]int, AN)
	for i, ele := range arrStr {
		arr[i], _ = strconv.Atoi(ele)
	}

	brr := make([]int, BN)
	for i, ele := range brrStr {
		brr[i], _ = strconv.Atoi(ele)
	}
	res := isSame(arr, brr, AN, BN)
	fmt.Println(res)
}
