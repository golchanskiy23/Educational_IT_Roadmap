/*
Given a string s, find the length of the longest substring without duplicate characters.
*/

func lengthOfLongestSubstring(s string) int {
    arr := []rune(s)
    m := make(map[rune]int)
    l,r, n, maximum := 0, 0, len(arr), 0
    for r < n{
        m[arr[r]] += 1
        for m[arr[r]] != 1{
            /*if m[arr[l]] == 1{
                delete(m, arr[l])
                break
            }*/
            m[arr[l]] -= 1
            l++
        }

        if r-l+1 > maximum{
            fmt.Print(l,r)
            maximum = r-l+1
        }
        r++
    }
    return maximum
}