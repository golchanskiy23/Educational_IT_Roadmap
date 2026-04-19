/*
Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.
*/

func longestSubarray(nums []int) int {
     prev, curr, best := 0, 0, 0

    for _, v := range nums {
        if v == 1 {
            curr++
        } else {
            if prev+curr > best {
                best = prev + curr
            }
            prev = curr
            curr = 0
        }
    }

    if prev+curr > best {
        best = prev + curr
    }

    if best == len(nums) {
        return best - 1
    }

    return best
}