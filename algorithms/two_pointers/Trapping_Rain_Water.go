/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
*/

func trap(nums []int) int {
    max, idx := -1,-1
    for i := 0; i < len(nums); i++{
        if nums[i] >= max{
            max = nums[i]
            idx = i
        }
    }

    stack, ans := 0,0
    for i := 0; i < idx; i++{
        if nums[i] > stack{
            stack = nums[i]
        } else{
            ans += (stack-nums[i])
        }
    }

    stack = 0
    for i := len(nums)-1; i > idx; i--{
        if nums[i] > stack{
            stack = nums[i]
        } else{
            ans += (stack-nums[i])
        }
    }
    
    return ans
}