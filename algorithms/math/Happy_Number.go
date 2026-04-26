/*
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.
*/

func sum(arr []int) int{
    s := 0
    for _, v := range arr {s+= (v*v)}
    return s
}

func split(n int) []int{
    arr := []int{}
    for n > 0{
        arr = append(arr, n%10)
        n /= 10
    }
    return arr
}

func isHappy(n int) bool {
    m := make(map[int]struct{})
    iter := 0
    for n != 1{
        if _, ok := m[n]; ok{
            return false
        } else{
            m[n] = struct{}{}
        }

        curr := sum(split(n))
        n = curr
        iter++
    }

    return true
}