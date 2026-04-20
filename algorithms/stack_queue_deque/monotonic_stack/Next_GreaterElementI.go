/*
The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.

You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.

For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater element of nums2[j] in nums2. If there is no next greater element, then the answer for this query is -1.

Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    stack := make([]int, 0)
    // num / next_greater
    m := make(map[int]int)
    for i := 0; i < len(nums2); i++{m[nums2[i]] = -1}

    for i := 0; i < len(nums2); i++{
        curr := nums2[i]
        if len(stack) == 0{
            stack = append(stack, curr)
            continue
        }

        for len(stack) > 0 && curr > stack[len(stack)-1]{
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            m[top] = curr
        }

        stack = append(stack, curr)
    }

    ans := make([]int, len(nums1))
    for i := 0; i < len(nums1); i++{
        ans[i] = m[nums1[i]]
    }

    return ans
}

func map_nextGreaterElement(nums1 []int, nums2 []int) []int {
    greater := make([]int, len(nums2))
    for i := 0; i < len(greater); i++{greater[i] = -1}

    // num / nex_greater
    m := make(map[int]int)
    for i := 0; i < len(nums2); i++{
        curr := nums2[i]
        m[curr] = -1
        for j := i+1; j < len(nums2); j++{
            if nums2[j] > curr{
                m[curr] = nums2[j]
                break
            }
        }
    }

    ans := make([]int, len(nums1))
    for i := 0; i < len(nums1); i++{
        ans[i] = m[nums1[i]]
    }

    return ans
}