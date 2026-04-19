/*
Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.
*/

func countSubstrings(s string) int {
    n,count := len(s),0
    for i := 0; i < n; i++{
        count += expand(s,i,i)
        count += expand(s,i,i+1)
    }
    return count
}

func expand(s string, l,r int) int{
    cnt := 0
    for l >= 0 && r < len(s) && s[l] == s[r]{
        cnt++
        l--
        r++
    }

    return cnt
}