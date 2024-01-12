package main

import (
	"fmt"
)

func fibo(N int) int {
	fmt.Println("N in", N)
	if N <= 1 {
		return N
	}

	return fibo(N-1) + fibo(N-2)
}

func mainFibo() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// N, _ := strconv.Atoi(scanner.Text())
	// fmt.Println("N", N)
	// fmt.Println("type of N", reflect.TypeOf(N))
	val := fibo(4)
	fmt.Println(val)
}
