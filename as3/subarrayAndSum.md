To demonstrate where the condition if i-modMap[mod] >= maxLength { modMap[mod] = i } is helpful, let's consider an example where an array and a value of K lead to potentially overlapping subarrays that meet our criteria, and see how updating modMap avoids counting these overlaps.

Example Array
Consider the array arr = [1, 2, 3, 4, 2, 1, 1] with K = 3.

First Loop - Finding Maximum Length
Assume from the first loop, we find that the maxLength of a subarray whose sum is not divisible by K is 4.

Second Loop - Counting Subarrays & Updating modMap
Now, in the second loop, we are counting subarrays of this maxLength and updating modMap:

First Subarray:

At index 3, we have a subarray [1, 2, 3, 4] with sum 10 and 10 % 3 = 1. This subarray's length is 4, which matches maxLength.
We increment our count to 1.
We update modMap[1] to 3.
Potential Overlap Without Updating modMap:

Now, suppose later in the array, say at index 6, we again have a sum modulo K of 1. This sum would be the result of the subarray [1, 2, 3, 4, 2, 1, 1].
If we hadn't updated modMap[1] to 3, it would still point to an earlier index (say 0). Then, we might incorrectly count a subarray that includes elements from our first subarray, leading to overlap.
Avoiding Overlap With Updated modMap:

Because we updated modMap[1] to 3, any new subarray ending with a sum % K of 1 will be considered starting from index 3 or later.
So, at index 6, when we see sum % K = 1 again, we look for a subarray starting from index 4 (next to where the last counted subarray ended) to index 6. This subarray is [2, 1, 1] with sum 4 and 4 % 3 = 1. It's a new, distinct subarray that doesn't overlap with [1, 2, 3, 4].
Conclusion
In this example, updating modMap[mod] prevents us from recounting subarrays that overlap with already counted subarrays. By ensuring each counted subarray starts from the most recent index for each modulo result, we accurately count distinct subarrays of maxLength whose sum is not divisible by K.
