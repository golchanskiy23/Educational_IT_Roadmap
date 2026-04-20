/*
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.
*/

func dailyTemperatures(temperatures []int) []int {
    ans := make([]int, len(temperatures))
    type pair struct{
        val, idx int
    }

    stack := make([]pair, 0)

    for i := 0; i < len(temperatures); i++{
        curr := temperatures[i]
        if len(stack) == 0{
            stack = append(stack, pair{
                val: curr,
                idx: i,
            })
            continue
        }

        for len(stack) > 0 && curr > stack[len(stack)-1].val {
            idx := stack[len(stack)-1].idx
            ans[idx] = i-idx
            stack = stack[:len(stack)-1]
        }
        
        stack = append(stack, pair{
            val: curr,
            idx: i,
        })
    }

    for _, p := range stack{
        ans[p.idx] = 0
    }

    return ans
}