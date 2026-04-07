/*
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
A subarray is a contiguous non-empty sequence of elements within an array.
*/

func subarraySum(nums []int, k int) int {
    if len(nums) == 1{
        if k == nums[0]{
            return 1
        }
        return 0
    }
	
    prefix := make([]int, len(nums)+1)
    m := make(map[int]int)
    cnt := 0
    m[0] = 1
    for i := 1; i <= len(nums); i++{
        prefix[i] = prefix[i-1]+nums[i-1]
        if freq , ok := m[prefix[i]-k]; ok{
            cnt += freq
        }
        m[prefix[i]]++
    }
    
    return cnt
}