package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
Given an array of integers, find the number of distinct pairs of a,b,c which Sum to 0.
Some notes, the pairs should be distinct</p><p>the pairs should be such that a<b<c
sort the list before printing, eg[ [-2,0,2],[-1,-1,2],[-1,0,1] ]

	// output
	// `3
	// -2 0 2
	// -1 -1 2
	// -1 0 1`
*/

// TC O(N^3), SC O(N^3)
func getTripletsBrute(arr []int, N int) [][]int {
	var triplets [][]int
	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					triplet := []int{arr[i], arr[j], arr[k]}
					triplets = append(triplets, triplet)
				}
			}
		}
	}
	return triplets
}

func removeDuplicates(triplets [][]int) [][]int {
	seen := make(map[[3]int]bool)
	var uniqueTriplets [][]int
	for _, triplet := range triplets {
		sort.Ints(triplet)
		t := [3]int{triplet[0], triplet[1], triplet[2]}
		if !seen[t] {
			seen[t] = true
			uniqueTriplets = append(uniqueTriplets, triplet)
		}
	}
	return uniqueTriplets
}

func findTriplets(arr []int, N int) [][]int {
	sort.Ints(arr)

	var triplets [][]int

	for i := 0; i < N-2; i++ {
		// Skip duplicate elements
		// L and R are still at the smae place
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		l, r := i+1, N-1
		for l < r {
			sum := arr[i] + arr[l] + arr[r]
			if sum == 0 {
				triplets = append(triplets, []int{arr[i], arr[l], arr[r]})
				l++
				r--
				// Skip duplicates for l and r
				// L and R are still at the smae place
				for l < r && arr[l] == arr[l-1] {
					l++
				}
				// L and R are still at the smae place
				for l < r && arr[r] == arr[r+1] {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return triplets
}

func mainFindTriplets() {

	input := `1
6
-1 0 1 2 -1 -2`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		NK := strings.Fields(scanner.Text())
		N, _ := strconv.Atoi(NK[0])

		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())
		arr := make([]int, N)
		for index, elem := range arrStr {
			arr[index], _ = strconv.Atoi(elem)
		}
		// triplets := getTriplets(arr, N)
		triplets := findTriplets(arr, N)
		// uniqTriplets := removeDuplicates(triplets)
		// fmt.Println("uniqTriplets", uniqTriplets)

		// fmt.Println(len(uniqTriplets))

		// sort.Slice(uniqTriplets, func(i, j int) bool {
		// 	if uniqTriplets[i][0] != uniqTriplets[j][0] {
		// 		return uniqTriplets[i][0] < uniqTriplets[j][0]
		// 	}
		// 	if uniqTriplets[i][1] != uniqTriplets[j][1] {
		// 		return uniqTriplets[i][1] < uniqTriplets[j][1]
		// 	}
		// 	return uniqTriplets[i][2] < uniqTriplets[j][2]
		// })

		for _, elem := range triplets {
			fmt.Println(elem[0], elem[1], elem[2])
		}
	}

}
