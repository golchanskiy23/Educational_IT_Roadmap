/*
Given a string s, return the longest palindromic substring in s.
*/

// Manacker algorithm
func longestPalindrome(s string) string {
    runes := []rune(s)
    n := len(runes)
    
    t := make([]rune, 2*n+1)
    for i := 0; i < n; i++ {
        t[2*i] = '#'
        t[2*i+1] = runes[i]
    }
    t[2*n] = '#'
    
    m := len(t)
    d := make([]int, m)
    l, r := 0, -1
    
    for i := 0; i < m; i++ {
        k := 0
        if i <= r {
            k = min(r-i, d[l+r-i])
        }
        
        for i+k+1 < m && i-k-1 >= 0 && t[i+k+1] == t[i-k-1] {
            k++
        }
        
        d[i] = k
        
        if i+k > r {
            l, r = i-k, i+k
        }
    }
    
    bestCenter, bestRadius := 0, 0
    for i, radius := range d {
        if radius > bestRadius {
            bestRadius = radius
            bestCenter = i
        }
    }
    
    start := (bestCenter - bestRadius) / 2
    end := (bestCenter + bestRadius) / 2
    return string(runes[start:end])
}