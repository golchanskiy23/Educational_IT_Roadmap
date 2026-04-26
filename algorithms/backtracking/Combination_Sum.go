/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
*/

func create(ans *[][]int, tmp *[]int, arr []int, target, idx int){
    if target == 0{
        cp := make([]int, len(*tmp))
        copy(cp, *tmp)
        *ans = append(*ans, cp)
        return
    }
    if target < 0 {return}

    for i := idx; i < len(arr); i++{
        *tmp = append(*tmp, arr[i])
        create(ans, tmp, arr, target-arr[i], i)
        *tmp = (*tmp)[:len(*tmp)-1]
    }
}

func combinationSum(candidates []int, target int) [][]int {
    var ans [][]int
    var tmp []int
    create(&ans, &tmp, candidates, target, 0)
    return ans
}