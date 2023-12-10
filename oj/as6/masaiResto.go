package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type StackR struct {
	Data []string
}

func NewStackR() StackR {
	return StackR{
		Data: []string{},
	}
}

func (s *StackR) push(val string) {
	s.Data = append(s.Data, val)
}

func (s *StackR) pop() {
	if len(s.Data) != 0 {
		lastElem := s.Data[len(s.Data)-1]
		fmt.Println(lastElem)
		s.Data = s.Data[:len(s.Data)-1]
	} else {
		fmt.Println("No Food")
	}
}

func main() {
	input := `6
1
2 5
2 7
2 9
1
1`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	stack := NewStackR()
	for i := 0; i < T; i++ {
		scanner.Scan()
		val := strings.Fields(scanner.Text())
		if len(val) > 1 {
			stack.push(val[1])
		} else {
			stack.pop()
		}
	}
}
