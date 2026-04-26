/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/

func merge(intervals [][]int) [][]int {
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
        return arr[i].l < arr[j].l
    })

    ans := make([][]int, 0)
    tmp := []int{}
    for i := 0; i < len(arr); i++{
        if len(tmp) == 0 {
            tmp = append(tmp, arr[i].l, arr[i].r)
            continue
        }

        if tmp[1] >= arr[i].l{
            if  tmp[1] < arr[i].r{
                tmp[1] = arr[i].r   
            }
        } else{
            ans = append(ans, tmp)
            tmp = []int{arr[i].l, arr[i].r}
        }
    }

    ans = append(ans, tmp)

    return ans
}