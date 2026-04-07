/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
*/

func isAnagram(s string, t string) bool {
    r1,r2 := []rune(s), []rune(t)
    sort.Slice(r1, func(i,j int) bool{
        return byte(r1[i]) < byte(r1[j])
    })
    sort.Slice(r2, func(i,j int) bool{
        return byte(r2[i]) < byte(r2[j])
    })

    if len(r1) != len(r2){return false}
    for i := 0; i < len(r1); i++{
        if r1[i] != r2[i]{
            return false
        }
    }
    return true
}