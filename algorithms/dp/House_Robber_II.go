/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
*/

func max(a,b int) int{
    if a > b{
        return a
    }

    return b
}

func util(nums []int) int{
    // найти максимум, не беря два подряд значения массива
    // то есть либо грабим этот дом, если ограбили позапрошлый,
    // либо грабим предыдущий не грабя текущий

    // max(dp[i-2]+nums[i], dp[i-1])
    // [1,2,3,1]
    // [0,1,2,4,3]


    // дома по кругу!!!
    dp := make([]int, len(nums)+1)
    dp[1] = nums[0]
    for i := 2; i <= len(nums); i++{
        dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
    }

    var ans int
    for _, v := range dp{
        ans = max(ans, v)
    }

    return ans
}

func rob(nums []int) int {
    n := len(nums)
    if len(nums) == 1{return nums[0]}
    return max(util(nums[:n-1]), util(nums[1:]))   
}