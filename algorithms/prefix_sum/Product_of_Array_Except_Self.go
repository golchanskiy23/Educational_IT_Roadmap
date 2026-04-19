/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/

func productExceptSelf(nums []int) []int {
    n := len(nums)
    prefix,suffix,arr := make([]int, n+1), make([]int, n+1), make([]int, n)
    prefix[0],suffix[n]=1,1

    for i := 1; i <= n; i++{
        prefix[i] = (prefix[i-1]*nums[i-1])
    }

    for i := n-1; i >= 0; i--{
        suffix[i] = (suffix[i+1]*nums[i])
    }

    for i := 0; i < n;i++{
        arr[i] = prefix[i]*suffix[i+1]
    }
    
    return arr
}