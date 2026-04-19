/*
Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the number of tuples (i, j, k, l) such that:

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    m12, m34 := make(map[int]int), make(map[int]int)
    n := len(nums1)
    for i := 0; i < n; i++{
        for j := 0; j < n; j++{
            m12[nums1[i]+nums2[j]]++
            m34[nums3[i]+nums4[j]]++
        }
    }

    ans := 0
    for k,v := range m12{
        ans += (v*m34[-k])
    }
    return ans
}