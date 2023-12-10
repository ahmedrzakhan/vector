package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// TC O(N^2), SC O(N^2)
func rotateMatrix(matrix [][]int, N int) [][]int {
	// Create a new matrix of the same size initialized with zeros
	rotated := make([][]int, N)
	for i := range rotated {
		rotated[i] = make([]int, N)
	}

	// The number of layers in the matrix
	layers := N / 2

	for layer := 0; layer < layers; layer++ {
		first, last := layer, N-layer-1

		// Top edge
		for i := first; i < last; i++ {
			rotated[first][i+1] = matrix[first][i]
		}

		// Right edge
		for i := first; i < last; i++ {
			rotated[i+1][last] = matrix[i][last]
		}

		// Bottom edge
		for i := first; i < last; i++ {
			rotated[last][i] = matrix[last][i+1]
		}

		// Left edge
		for i := first; i < last; i++ {
			rotated[i][first] = matrix[i+1][first]
		}
	}

	// For odd dimension matrices, set the center element
	if N%2 != 0 {
		center := layers
		rotated[center][center] = matrix[center][center]
	}

	return rotated
}

func mainRot() {
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

	rotated := rotateMatrix(matrix, N)

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
