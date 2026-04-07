/*
Given a string s, return the number of unique palindromes of length three that are a subsequence of s.
Note that even if there are multiple ways to obtain the same subsequence, it is still only counted once.
A palindrome is a string that reads the same forwards and backwards.
A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
For example, "ace" is a subsequence of "abcde".
*/

func countPalindromicSubsequence(s string) int {
    ids := make(map[rune][2]int)
    runes := []rune(s)
    prefix := make([]int, len(runes)+1)
    for i, r := range runes{
        if _, ok := ids[r]; !ok{
            ids[r] = [2]int{i,i}
        }
        val := ids[r]
        val[1] = i
        ids[r] = val
        prefix[i+1] = len(ids)
    }
    // [0,1,1,2,3,3]
    ans := 0
    for _, v := range ids{
        if v[0] == v[1]{
            continue
        }
        seen := make(map[rune]struct{})
        for i := v[0]+1; i < v[1]; i++{
            seen[runes[i]] = struct{}{}
        }
        ans += len(seen)
    }
    return ans
}