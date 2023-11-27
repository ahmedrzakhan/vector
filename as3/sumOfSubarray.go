package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
You are given an array of N integers and a single integer K.
You need to find out if there is a subarray, which has the sum exactly as K.
*/

// 5 3
// 1 2 1 3 4
// Yes

// 4 5
// 1 2 1 3
// No

// 3 2
// 1 2 1
// Yes

// hasSubarrayWithSum checks if there is a subarray with sum equal to K
/**
TC - O(N^2), SC - O(N)
*/
func hasSubarrayWithSumBr(arr []int, N int, K int) string {
	// Iterate over all subarrays
	for i := 0; i < N; i++ {
		sum := 0
		for j := i; j < N; j++ {
			sum += arr[j]
			// Check if the current subarray's sum equals K
			if sum == K {
				return "Yes"
			} else if sum > K {
				// reset sum to 0 if sum is grater than K
				sum = 0
			}
		}
	}

	return "No"
}

/*
*
Checking for Subarray with Sum K:

At each element, the function checks if sum - K exists in the HashMap sums.
Why sum - K? If sum is the current total sum and there was a previous prefix sum
equal to sum - K, then the elements between that prefix and the current element
add up to K.
For example, if sum is 10 and K is 3, you look for a prefix sum of 7. If you find
it, it means the elements between this prefix and the current index sum up to 3.
If such a previous prefix sum exists, it means there is a subarray that sums to K,
and the function returns "Yes".
*/

/*
*

The sum - K expression is derived from the logic we use to find a subarray with a sum equal
to K. Let's break down how we arrive at this expression:

Concept of Prefix Sum:

A prefix sum is the sum of elements from the start of the array up to a given index.
For example, in an array [3, 1, 4, 2], the prefix sum up to the third element (4) is 3 + 1 + 4 = 8.
Running Sum (sum) in the Function:

As the function iterates through the array, sum keeps track of the sum of elements from the start
to the current element.
This running sum is essentially the prefix sum at each index of the array.
Finding Subarrays that Sum to K:

Suppose we are at a certain point in the array where the running sum is sum. We want to know if
there's a contiguous subarray that ends at this point with a total sum of K.
If such a subarray exists, then the sum of the elements before this subarray would be sum - K. This
is because the elements of the subarray itself add up to K, and the total sum up to this point is sum.
Rearranging the Equation:

Let's assume that the sum of the subarray we are looking for is K.
If the running sum at a certain point is sum, and there is a subarray ending at this point that sums

	to K, then the sum of the array before the start of this subarray must be sum - K.

Mathematically, it's like saying:
Total Sum (up to current index) = Sum of Subarray (equals K) + Sum before Subarray.
sum = K + (sum - K).
Checking for sum - K:

The function checks if sum - K exists in the hashmap of prefix sums.
If sum - K is found in the hashmap, it means there was a point in the array where the sum of elements
before that point was sum - K. And from that point to the current point, the sum of elements is K.
Example:

Consider an array [3, 1, 4, 2] and K = 5.
When we reach the third element (4), our running sum (sum) is 8.
We check if sum - K (which is 8 - 5 = 3) is in our hashmap.
If 3 is in the hashmap, it means there is a subarray before this point that sums up to 3, and the elements
from there to our current element sum up to 5.
*/
/**
TC - O(N), SC - O(N)
*/
func hasSubarrayWithSum(arr []int, N int, K int) string {
	sum := 0
	sums := make(map[int]struct{}) // A hash map to store the prefix sums

	// A zero sum is always present (empty subarray)
	sums[0] = struct{}{}

	for _, val := range arr {
		sum += val
		if _, ok := sums[sum-K]; ok {
			return "Yes"
		}
		// Store the running sum in the hash map
		sums[sum] = struct{}{}
	}

	return "No"
}

// main function to test hasSubarrayWithSum
func main() {

	input := `3
5 3
1 2 1 3 4
4 5
1 2 1 3
3 2
1 2 1`

	// scanner := bufio.NewScanner(os.Stdin)
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		NK := strings.Fields(scanner.Text())
		N, _ := strconv.Atoi(NK[0])
		K, _ := strconv.Atoi(NK[1])

		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())
		arr := make([]int, N)
		for j, s := range arrStr {
			arr[j], _ = strconv.Atoi(s)
		}
		fmt.Println(hasSubarrayWithSum(arr, N, K))
	}

}
