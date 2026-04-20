/*
Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.
*/

func max(a,b int) int{
    if a > b{
        return a
    }

    return b
}

func largestRectangleArea(heights []int) int {
    ans := 0
    type pair struct{
        val, idx int
    }

    stack := make([]pair, 0)

    for i := 0; i < len(heights); i++{
        curr, start := heights[i], i
        if len(stack) == 0{
            stack = append(stack, pair{
                val: curr,
                idx: i,
            })
            continue
        }

        for len(stack) > 0 && curr < stack[len(stack)-1].val {
            idx := stack[len(stack)-1].idx
            ans = max((i-idx)*stack[len(stack)-1].val , ans)
            stack = stack[:len(stack)-1]
            start = idx
        }
        
        stack = append(stack, pair{
            val: curr,
            idx: start,
        })
    }

    for len(stack) > 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ans = max(ans, (len(heights)-top.idx)*top.val)
    }

    return ans
}