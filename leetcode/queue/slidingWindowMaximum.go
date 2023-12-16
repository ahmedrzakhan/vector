package main

import "fmt"

/**
239. Sliding Window Maximum

You are given an array of integers nums, there is a sliding window of size k which is moving
from the very left of the array to the very right. You can only see the k numbers in the window.
Each time the sliding window moves right by one position.

Return the max sliding window.

Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Example 2:

Input: nums = [1], k = 1
Output: [1]

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
*/

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	result := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		// Remove front elements that are out of the current window
		if len(deque) > 0 && deque[0] == i-k {
			deque = deque[1:]
		}

		// Remove elements smaller than the current element from the back
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Add the current element index
		deque = append(deque, i)

		// Add the front element of deque to result if the first window is done
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func mainWin() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k)) // Output: [3, 3, 5, 5, 6, 7]
}
