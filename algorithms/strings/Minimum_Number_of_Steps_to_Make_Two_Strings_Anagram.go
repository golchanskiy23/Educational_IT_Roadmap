/*
You are given two strings of the same length s and t. In one step you can choose any character of t and replace it with another character.

Return the minimum number of steps to make t an anagram of s.

An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.
*/

func min(a,b int) int{
    if a < b{return a}
    return b
}

func minSteps(s string, t string) int {
    sr, tr := []rune(s), []rune(t)

    sm, tm := make(map[rune]int), make(map[rune]int)
    diff := 0
    for i := 0; i < len(sr); i++{
        sm[sr[i]]++
        tm[tr[i]]++
    }

    for k,v := range sm{
        if val, ok := tm[k]; ok{
            diff += min(v,val)
        }
    }
    return len(sr)-diff
}