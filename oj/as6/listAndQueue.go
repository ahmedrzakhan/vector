package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Queue struct {
	Data []string
}

func NewQueue() Queue {
	return Queue{
		Data: []string{},
	}
}

func (s *Queue) enqueue(val string) {
	s.Data = append(s.Data, val)
}

func (s *Queue) dequeue() {
	if len(s.Data) == 0 {
		fmt.Println("Empty")
	} else {
		fronElement := s.Data[0]
		s.Data = s.Data[1:]
		fmt.Println(fronElement)
	}
}

func mainL() {
	input := `5
E 2
E 3
D
D
D`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	queue := NewQueue()
	for i := 0; i < T; i++ {
		scanner.Scan()
		val := strings.Fields(scanner.Text())
		if len(val) > 1 {
			queue.enqueue(val[1])
		} else if val[0] == "D" {
			queue.dequeue()
		}
	}

}
