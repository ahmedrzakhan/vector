package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func mainProcessMatrix() {
	input := `4
1 2 3 4
1 2 3 4
1 2 3 4
1 2 3 4`
	// `1 1 2 3
	// 1 2 2 4
	// 1 3 3 4
	// 2 3 4 4`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		rowStr := scanner.Text()
		rowElements := strings.Fields(rowStr)
		matrix[i] = make([]int, N)
		for j, elemStr := range rowElements {
			matrix[i][j], _ = strconv.Atoi(elemStr)
		}
	}

	rotated := matrix

	for _, row := range rotated {
		for j, num := range row {
			fmt.Print(num)
			if j < N-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
