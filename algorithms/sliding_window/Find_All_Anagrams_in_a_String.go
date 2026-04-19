/*
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order
*/

func equal(a,b []int) bool{
    for i := 0; i < 26; i++{
        if a[i] != b[i]{
            return false
        }
    }
    return true
}

func findAnagrams(s string, p string) []int {
    if len(s) < len(p){
        return []int{}
    }
    sCount, pCount := make([]int, 26), make([]int, 26)
    for _, v := range p{
        pCount[v-'a']++
    }

    l,r := 0,0
    ans := make([]int, 0)
    for r < len(s){
        sCount[s[r]-'a']++

        if r-l+1 > len(p) {
            sCount[s[l]-'a']--
            l++
        }

        if r-l+1 == len(p) && equal(pCount, sCount) {
            ans = append(ans, l)
        }
        r++
    }
    return ans
}

/*func slow_findAnagrams(s string, p string) []int {
    pattern := []rune(p)
    str := []rune(s)
    if len(str) < len(pattern){
        return []int{}
    }
    ans := make([]int, 0)
    l,r := 0, len(pattern)
    
    slices.Sort(pattern)
    
    for r <= len(str){
        curr := make([]rune, len(pattern))
        copy(curr, str[l:r])

        slices.Sort(curr)
        if slices.Equal(curr, pattern){
            ans = append(ans, l)
        }
        l++
        r++
    }
    return ans
}*/