/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.
*/

func min(a,b int) int{
    if a < b{
        return a
    }
    return b
}

// [2,3,1,2,4,3]  1 4 4
func minSubArrayLen(target int, nums []int) int {
    ans, sum, n := math.MaxInt, 0, len(nums)
    l,r := 0,0
    for r < n{
        sum += nums[r]
        for sum >= target{
            ans = min(ans, r-l+1)
            sum -= nums[l]
            l++
        }
        r++
    }

    if ans == math.MaxInt{
        return 0
    }
    return ans
}