/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/

func create(ans *[][]int, tmp *[]int, nums []int, used []bool){
    if len(*tmp) == len(nums){
        cp := make([]int, len(*tmp))
        copy(cp, *tmp)
        *ans = append(*ans, cp)
        return
    }

    for i := 0; i < len(nums); i++{
        if used[i] {continue}
        used[i] = true
        *tmp = append(*tmp, nums[i])
        create(ans, tmp, nums, used)
        *tmp = (*tmp)[:len(*tmp)-1]
        used[i] = false
    }
}

func permute(nums []int) [][]int {
    var ans [][]int
    var tmp []int
    used := make([]bool, len(nums))
    create(&ans, &tmp, nums, used)
    return ans
}