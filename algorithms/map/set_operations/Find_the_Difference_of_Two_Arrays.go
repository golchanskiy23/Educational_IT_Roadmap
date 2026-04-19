/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.
*/

func findDifference(nums1 []int, nums2 []int) [][]int {
    ans := make([][]int, 2)
    m1, m2 := make(map[int]struct{}) , make(map[int]struct{})
    for i := 0; i < len(nums1); i++ {m1[nums1[i]] = struct{}{}}
    for i := 0; i < len(nums2); i++ {m2[nums2[i]] = struct{}{}}
    for k,_ := range m1{
        if _, ok := m2[k]; !ok{
            ans[0] = append(ans[0], k)
        }
    }

    for k, _ := range m2{
        if _, ok := m1[k]; !ok{
            ans[1] = append(ans[1], k)
        }
    }

    return ans
}