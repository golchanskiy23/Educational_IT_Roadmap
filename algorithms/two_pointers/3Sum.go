/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/

func threeSum(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)

    ans := make([][]int, 0)
    for i := 0; i < n; i++{
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        
        l,r := i+1, n-1
        for l < r{
            curr := nums[i]+nums[l]+nums[r] 
            if curr == 0{
                a := []int{nums[i],nums[l],nums[r]}
                ans = append(ans, a)
                for l < r && nums[l] == nums[l+1] {
                    l++
                }
                
                for l < r && nums[r] == nums[r-1] {
                    r--
                }

                l++
                r--
            } else if curr > 0{
                r--
            } else{
                l++
            }
        }
    }
    return ans
}