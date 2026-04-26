/*
Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note that intervals which only touch at a point are non-overlapping. For example, [1, 2] and [2, 3] are non-overlapping.
*/

func eraseOverlapIntervals(intervals [][]int) int {
    type interval struct{
        l,r int
    }
    arr := make([]interval, len(intervals))
    for i := 0; i < len(intervals); i++{
        arr[i] = interval{
            l: intervals[i][0],
            r: intervals[i][1],
        }
    }
    
    sort.Slice(arr, func(i,j int)bool{
        if arr[i].l == arr[j].l{
            return arr[i].r < arr[j].r
        }
        return arr[i].l < arr[j].l
    })

    ans := 0
    tmp := []int{}
    for i := 0; i < len(arr); i++{
        if len(tmp) == 0 {
            tmp = append(tmp, arr[i].l, arr[i].r)
            continue
        }

        if tmp[1] > arr[i].l{
            fmt.Println(tmp)
            tmp[1] = min(tmp[1], arr[i].r)
            ans++
        } else{
            tmp = []int{arr[i].l, arr[i].r}
        }
    }
    return ans
}