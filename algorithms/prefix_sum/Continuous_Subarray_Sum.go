/*
Given an integer array nums and an integer k, return true if nums has a good subarray or false otherwise.

A good subarray is a subarray where:

its length is at least two, and
the sum of the elements of the subarray is a multiple of k.
Note that:

A subarray is a contiguous part of the array.
An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.
*/

func checkSubarraySum(nums []int, k int) bool {
    if len(nums) == 1{
        if nums[0] == 0 || (nums[0] == k){
            return false
        }
        return true
    }

    m := make(map[int]int)
    sum := 0
    m[0] = -1
    for i := 0; i < len(nums); i++{
        sum += nums[i]
        rem := (sum)%k
        if val, ok := m[rem]; ok{
            if i-val > 1{
                fmt.Println(i,val, sum)
                return true
            } 
        } else{
            m[rem] = i
        }
    }

    return false
}