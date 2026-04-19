/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.
*/

func min(a,b int) int{
    if a < b{
        return a
    }
    return b
}

func intersect(nums1 []int, nums2 []int) []int {
    ans := make([]int, 0)
    m1, m2 := make(map[int]int) , make(map[int]int)
    for i := 0; i < len(nums1); i++ {m1[nums1[i]]++}
    for i := 0; i < len(nums2); i++ {m2[nums2[i]]++}
    for k,v := range m1{
        if m2[k] != 0{
            m := min(v, m2[k])
            for m > 0{
                ans = append(ans, k)
                m--
            }
        }
    }

    return ans
}