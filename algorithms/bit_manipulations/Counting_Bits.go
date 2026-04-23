/*
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.
*/

func countBits(n int) []int {
    if n == 0{return []int{0}}
    if n == 1{return []int{0,1}}

    ans := make([]int, n+1)
    ans[0], ans[1] = 0,1
    upd, dist := 0, 1
    for i := 2; i <= n; i++{
        ans[i] = ans[i-dist]+upd
        upd++
        if upd == 2{
            upd = 0
        } else{
            dist++
        }
    }

    return ans
}