/*
Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
*/

type NumArray struct {
    prefix []int
}


func Constructor(nums []int) NumArray {
    curr := make([]int, len(nums)+1)
    for i := 1; i <= len(nums); i++{
        curr[i] = (curr[i-1]+nums[i-1])
    }
    fmt.Println(curr)
    return NumArray{
        prefix: curr,
    }
}

func (this *NumArray) SumRange(l int, r int) int {
    return this.prefix[r+1]-this.prefix[l]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */