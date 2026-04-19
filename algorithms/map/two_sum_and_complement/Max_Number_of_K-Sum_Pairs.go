/*
You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.
*/

func maxOperations(nums []int, k int) int {
    m := make(map[int]int)

    for _, v := range nums {
        m[v]++
    }

    ans := 0

    for x, cnt := range m {
        if cnt == 0 {
            continue
        }

        y := k - x

        if x == y {
            ans += cnt / 2
            m[x] = 0
        } else if m[y] > 0 {
            pairs := min(cnt, m[y])
            ans += pairs
            m[x] -= pairs
            m[y] -= pairs
        }
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func sort_maxOperations(nums []int, k int) int {
    sort.Ints(nums)
    l,r,ans := 0, len(nums)-1,0
    for l < r{
        curr := nums[l]+nums[r]
        if curr == k{
            ans++
            l++
            r--
        } else if curr > k{
            r--
        } else{
            l++
        }
    }
    return ans
}