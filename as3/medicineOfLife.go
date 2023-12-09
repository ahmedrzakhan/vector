package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
given array in ascending sorted order, combination of two adds up to K
print Possible, else Not Possible
*/

// TC - O(N) SC - O(1)
func isPossible(arr []int, N, K int) string {
	left, right := 0, N-1
	for left < right {
		if arr[left]+arr[right] == K {
			return "Possible"
		}

		if arr[left]+arr[right] < K {
			left++
		}

		if arr[left]+arr[right] > K {
			right--
		}
	}
	return "Impossible"
}

func mainMed() {
	input := `2
	5 7
	1 2 3 4 5
	5 12
	1 2 3 4 5`
	// 	Possible
	// Impossible
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
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
		for index, ele := range arrStr {
			arr[index], _ = strconv.Atoi(ele)
		}
		poss := isPossible(arr, N, K)
		fmt.Println(poss)
	}

}
