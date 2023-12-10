package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Marks int
}

// TC O(N^2) SC O(1)
func bubbleSort(students []Student, N int) []Student {
	for i := 0; i < N; i++ {
		for j := 0; j < N-i-1; j++ {
			if students[j].Marks < students[j+1].Marks ||
				students[j].Marks == students[j+1].Marks &&
					students[j].Name > students[j+1].Name {
				students[j], students[j+1] = students[j+1], students[j]
			}
		}
	}
	return students
}

func mainLeaderboard() {
	input := `6
	rancho 45
	chatur 32
	raju 30
	farhan 28
	virus 32
	joy 45`
	// output
	// 	1 joy
	// 1 rancho
	// 3 chatur
	// 3 virus
	// 5 raju
	// 6 farhan

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	students := make([]Student, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		marks, _ := strconv.Atoi(fields[1])
		students[i] = Student{Name: fields[0], Marks: marks}
	}

	bubbleSort(students, N)
	rank := 1
	for i, student := range students {
		if i > 0 && student.Marks != students[i-1].Marks {
			rank = i + 1
		}
		fmt.Printf("%d %s\n", rank, student.Name)
	}
}
