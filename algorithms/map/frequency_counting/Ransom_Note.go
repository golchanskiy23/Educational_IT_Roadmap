/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.
*/

func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[rune]int)
    runes := []rune(magazine)
    for i := 0; i < len(runes); i++{
        m[runes[i]]++
    }

    s := []rune(ransomNote)
    for i := 0; i < len(s); i++{
        if m[s[i]] == 0{
            return false
        }
        m[s[i]]--
    }
    
    return true
}