/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int{
    m := make(map[int]int)
    for i := 0; i < len(nums); i++{
        if val, ok := m[target-nums[i]]; ok{
            if 2*nums[i] == target && val != i{
                return []int{val, i}
            }
        } else{
            m[nums[i]] = i
        }
    }

    for i := 0; i < len(nums); i++{
        if idx, ok := m[target-nums[i]]; ok{
            if idx != i{
                return []int{idx, i}
            }
        }
    }

    return []int{}
}

func twop_sort_twoSum(nums []int, target int) []int{
    type pair struct{
        val,idx int
    }

    arr := make([]pair, len(nums))
    for i := 0; i < len(nums); i++{
        arr[i] = pair{
            val: nums[i],
            idx: i,
        }
    }

    sort.Slice(arr, func(i,j int)bool{
        return arr[i].val < arr[j].val
    })

    l,r := 0, len(nums)-1
    for l < r{
        if arr[l].val+arr[r].val == target{
            return []int{arr[l].idx, arr[r].idx}
        } else if arr[l].val+arr[r].val > target{
            r--
        } else{
            l++
        }
    }
    return []int{}
}