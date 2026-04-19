/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
*/

func min(a,b int) int{
    if a < b{
        return a
    }
    return b
}

func findMaxAverage(nums []int, k int) float64 {
    max, sum, n := 0.0, 0.0, len(nums)
    curr := min(k,n)
    for i := 0; i < curr; i++{
        sum += float64(nums[i])
    }
    max = sum/float64(curr)
    l,r := 0,k-1
    for r < n{
        sum -= float64(nums[l])
        l++
        r++
        if r >= n{break}
        sum += float64(nums[r])
        if sum/float64(curr) > max{
            max = sum/float64(curr)
        }
    }
    return max
}