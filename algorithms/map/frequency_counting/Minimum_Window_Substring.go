/*
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.
*/

func min(a,b int) int{
    if a < b{
        return a
    }
    return b
}

func minWindow(s string, t string) string {
    if len(t) > len(s){
        return ""
    }

    need := make(map[byte]int)
    for i := 0; i < len(t); i++{
        need[t[i]]++
    }

    have := 0
    window := make(map[byte]int)

    l,r := 0,0
    minimum := math.MaxInt
    var ans string
    for r < len(s){
        c := s[r]
        if val, ok := need[c]; ok{
            window[c]++
            if val == window[c]{
                have++
            }
        }

        for have == len(need) {
            if r-l+1 < minimum {
                minimum = r - l + 1
                ans = s[l : r+1]
            }

            left := s[l]
            window[left]--

            if need[left] > 0 && window[left] < need[left] {
                have--
            }

            l++
        }
        
        r++
    }

    return ans
}