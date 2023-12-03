package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
input array 1 to N, no duplicates, one integermissing
return that
*/

// TC - O(NLogN) SC - O(1)
func getInteger(arr []int) int {
	N := len(arr)
	sort.Ints(arr)
	for i := 1; i < N; i++ {
		if arr[i-1] != i {
			return i
		}
	}
	return -1
}

func mainMissing() {
	input := `4 5 1 2`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1064*1064)
	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())
	var arr []int
	for _, ele := range arrStr {
		item, _ := strconv.Atoi(ele)
		arr = append(arr, item)
	}
	integer := getInteger(arr)
	fmt.Println(integer)
}
