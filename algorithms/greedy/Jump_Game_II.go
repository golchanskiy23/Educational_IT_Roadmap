/*
You are given a 0-indexed array of integers nums of length n. You are initially positioned at index 0.

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at index i, you can jump to any index (i + j) where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach index n - 1. The test cases are generated such that you can reach index n - 1.
*/

func jump(nums []int) int {
    jumps, currEnd, right := 0,0,0
    for i := 0; i < len(nums)-1; i++{
        right = max(right, i+nums[i])
        if i == currEnd{
            jumps++
            currEnd  = right
        }
    }

    return jumps
}