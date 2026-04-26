/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.
*/

func create(ans *[][]int, tmp *[]int, nums []int, idx int){
    cp := make([]int, len(*tmp))
    copy(cp, *tmp)
    *ans = append(*ans, cp)

    for i := idx; i < len(nums); i++{
        *tmp = append(*tmp, nums[i])
        create(ans, tmp, nums, i+1)
        *tmp = (*tmp)[:len(*tmp)-1]
    }
}

func subsets(nums []int) [][]int {
    var ans [][]int
    var tmp []int
    create(&ans, &tmp, nums, 0)
    return ans
}