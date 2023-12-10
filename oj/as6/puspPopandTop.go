package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	data []string
}

func NewStack() Stack {
	return Stack{
		data: []string{},
	}
}

func (s *Stack) push(val string) {
	s.data = append(s.data, val)
}

func (s *Stack) pop() {
	if len(s.data) != 0 {
		s.data = s.data[:len(s.data)-1]
	}
}

func (s *Stack) printTop() {
	if len(s.data) == 0 {
		fmt.Println("Empty!")
	} else {
		fmt.Println(s.data[len(s.data)-1])
	}
}

func mainS() {
	input := `6
2
2
2
2
1 4
3`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	stack := NewStack()

	for i := 0; i < T; i++ {
		scanner.Scan()
		val := strings.Fields(scanner.Text())

		if len(val) > 1 {
			stack.push(val[1])
		} else if val[0] == "2" {
			stack.pop()
		} else if val[0] == "3" {
			stack.printTop()
		}
	}
}
