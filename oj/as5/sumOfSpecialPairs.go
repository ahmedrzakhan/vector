package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
Write a program to find the sum of the absolute difference between all such pairs (A[i], A[j]) such that i < j and ( j-i ) is prime
*/

/**
NOTE:
every prime number (except 2 and 3) can be expressed in the form of 6k - 1 or 6k + 1, where k is an integer.
Increment by 6:
The loop starts with i = 5 and then increments i by 6 in each iteration (i += 6). This increment is based on the observation that every prime number (except 2 and 3) can be expressed in the form of 6k - 1 or 6k + 1, where k is an integer.

Why 6k - 1 and 6k + 1? Every integer can be expressed as either 6k, 6k + 1, 6k + 2, 6k + 3, 6k + 4, or 6k + 5. Among these:
6k, 6k + 2, and 6k + 4 are multiples of 2 (and hence not prime, except for 2 itself).
6k and 6k + 3 are multiples of 3 (and hence not prime, except for 3 itself).
That leaves 6k + 1 and 6k + 5. Note that 6k + 5 can be rewritten as 6(k + 1) - 1, which is of the form 6k - 1.
Conditions N % i == 0 || N % (i + 2) == 0:
This part of the loop is checking whether the number N is divisible by either i or i + 2.

N % i == 0: This checks if N is divisible by i, which takes on values 6k - 1 starting with 5, 11, 17, etc. If N is divisible by any such i, it means N is not a prime number.
N % (i + 2) == 0: Since i is of the form 6k - 1, i + 2 will be of the form 6k + 1. This condition checks if N is divisible by such numbers, starting with 7, 13, 19, etc. If N is divisible by any such i + 2, it also means N is not a prime number.

*/

func absInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

// O(N)
// isPrime checks if a number is prime in an unoptimized way
// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false // Numbers less than or equal to 1 are not prime
// 	}

// 	// Check divisibility by all numbers from 2 to n-1
// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false // If divisible, n is not prime
// 		}
// 	}

// 	return true // If no divisors found, n is prime
// }

// O(sqrt(n))
func isPrime(N int) bool {
	if N <= 1 {
		return false // Numbers less than or equal to 1 are not prime
	}
	if N <= 3 {
		return true // 2 and 3 are prime
	}

	// Check if n is divisible by 2 or 3
	if N%2 == 0 || N%3 == 0 {
		return false
	}

	// Check for higher divisors
	// i * i = N or i <= int(math.Sqrt(float64(N)))
	for i := 5; i*i <= N; i += 6 {
		if N%i == 0 || N%(i+2) == 0 {
			return false
		}
	}
	return true

}

// TC -O(N^2) SC -O(1)
func getSum(arr []int, N int) int {
	sum := 0
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			if i < j && isPrime(j-i) {
				sum = sum + absInt(arr[i]-arr[j])
			}
		}
	}
	return sum
}

func mainSumof() {
	input := `6
	1 2 3 5 7 12`
	// 45

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()

	arrStr := strings.Fields(scanner.Text())

	arr := make([]int, N)
	for index, elem := range arrStr {
		arr[index], _ = strconv.Atoi(elem)
	}

	sum := getSum(arr, N)
	fmt.Println(sum)
}
