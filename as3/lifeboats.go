package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
N passengers, boat has capcity of K
max 2 people can occupy a boat
tell number of boats required
*/

// TC - NLogN SC - O(1)
// func lifeboat(arr []int, N int, K int) int {
// 	sort.Ints(arr)
// 	count := 0

// 	for i := 0; i < N; i++ {
// 		if i != N-1 && arr[i]+arr[i+1] <= K {
// 			count++
// 			i++
// 		} else {
// 			count++
// 		}
// 	}
// 	return count
// }

func lifeboat(arr []int, N int, K int) int {
	count := 0
	left, right := 0, N-1
	sort.Ints(arr)

	for left <= right {
		if arr[left]+arr[right] <= K {
			left++
		}
		count++
		right--
	}

	return count
}

func mainLifeboat() {
	input := `2
4 5
3 5 3 4
4 3
1 2 2 3`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()

	T, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < T; i++ {
		scanner.Scan()
		NK := strings.Fields(scanner.Text())
		N, _ := strconv.Atoi(NK[0])
		K, _ := strconv.Atoi(NK[1])
		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())
		arr := make([]int, N)
		for i, ele := range arrStr {
			arr[i], _ = strconv.Atoi(ele)
		}
		resp := lifeboat(arr, N, K)
		fmt.Println(resp)
	}

}
