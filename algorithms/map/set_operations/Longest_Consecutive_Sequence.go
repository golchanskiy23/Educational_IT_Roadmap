/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.
*/

func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func longestConsecutive(nums []int) int{
    if len(nums) <= 1{
        return len(nums)
    }

    set := make(map[int]struct{})
    for i := 0; i < len(nums); i++{
        set[nums[i]] = struct{}{}
    }

    ans := 1
    for k, _ := range set{
        val := k
        if _, ok := set[val-1]; !ok{
            curr := 1
            for{
                if _, ok := set[val+1]; ok{
                    curr++
                    val++
                } else{
                    break
                }
            }
            ans = max(ans, curr)
        }
    }

    return ans
}

func set_longestConsecutive(nums []int) int {
    if len(nums) <= 1{
        return len(nums)
    }

    set := make(map[int]struct{})
    for i := 0; i < len(nums); i++{
        set[nums[i]] = struct{}{}
    }

    arr := make([]int, len(set))
    i := 0
    for k, _ := range set {
        arr[i] = k
        i++
    }
    sort.Ints(arr)

    ans, curr := 1,1
    for i := 1; i < len(arr); i++{
        if arr[i]-arr[i-1] == 1{
            curr++
            ans = max(ans, curr)
        } else{
            curr = 1
        }
    }

    return ans
}