package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
New Year Celebration

Description

New Year Celebration is coming near and DG is hosting a party for it. To make sure the party goes well she wants order in the party. There are two counters, one for ice-cream and the other for Cold-drinks. On the ice-cream counter, the line was formed in the form of a Queue and on the drinks counter, the line was formed in order of Stacks. She gave the management of the party to NV, NV made a plan to ask the manager following types of query.

Query types:

X : Number X enter the Queue.
X : Number X enter the Stack.
Output whoever is in front of the queue.
Output whoever is on top of the stack.
Which person is in front of the queue gets out of the queue and enters the stack.
Note: If the Queue or Stack is empty then output "-1".
Input

Input Format:

The first line of input will contain Q, which is the number of queries.
The next Q lines will be different queries based upon query types given.

Constraints:

1 <= Q <= 10^5
1 <= X <= 10^9

Output

Output will consist of integers based upon Query types. If Query type is 3 then Output whoever is in front of the queue, if Query type is 4 Output whoever is on top of the stack.

Sample Input 1

7
1 4
2 3
1 2
3
4
5
4

Sample Output 1

4
3
4
*/

type DataStructure struct {
	Queue QueueNewYear
	Stack StackNewYear
}

func newDataStructure() DataStructure {
	return DataStructure{
		Queue: newQueueNewYear(),
		Stack: newStackNewYear(),
	}
}

func (s *DataStructure) moveFromQUeueToStack() {
	elem := s.Queue.dequeue()
	s.Stack.push(elem)
}

type QueueNewYear struct {
	Data []string
}

func newQueueNewYear() QueueNewYear {
	return QueueNewYear{
		Data: []string{},
	}
}

func (s *QueueNewYear) enqueue(val string) {
	s.Data = append(s.Data, val)
}

func (s *QueueNewYear) peek() string {
	if len(s.Data) == 0 {
		return "-1"
	} else {
		firstElement := s.Data[0]
		// s.Data = s.Data[1:]
		return firstElement
	}
}

func (s *QueueNewYear) dequeue() string {
	if len(s.Data) == 0 {
		return "-1"
	} else {
		firstElement := s.Data[0]
		s.Data = s.Data[1:]
		return firstElement
	}
}

type StackNewYear struct {
	Data []string
}

func newStackNewYear() StackNewYear {
	return StackNewYear{
		Data: []string{},
	}
}

func (s *StackNewYear) push(val string) {
	s.Data = append(s.Data, val)
}

func (s *StackNewYear) pop() string {
	if len(s.Data) == 0 {
		return "-1"
	} else {
		lastElem := s.Data[len(s.Data)-1]
		// s.Data = s.Data[:len(s.Data)-1]
		return lastElem
	}
}

func mainNew() {
	input := `7
	1 4
	2 3
	1 2
	3
	4
	5
	4`
	// 	input := `10
	// 1 99999
	// 2 10000
	// 3
	// 4
	// 1 300
	// 4
	// 2 444
	// 5
	// 4
	// 3`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()

	T, _ := strconv.Atoi(scanner.Text())

	dataStructure := newDataStructure()

	for i := 0; i < T; i++ {
		scanner.Scan()
		item := strings.Fields(scanner.Text())
		if len(item) > 1 {
			if item[0] == "1" {
				dataStructure.Queue.enqueue(item[1])
			} else if item[0] == "2" {
				dataStructure.Stack.push(item[1])
			}
		} else {
			if item[0] == "3" {
				fmt.Println(dataStructure.Queue.peek())
			} else if item[0] == "4" {
				fmt.Println(dataStructure.Stack.pop())
			} else if item[0] == "5" {
				dataStructure.moveFromQUeueToStack()
			}
		}
	}
}
