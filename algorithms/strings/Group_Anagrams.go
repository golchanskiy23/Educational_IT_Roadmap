/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
*/

func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    for _, val := range strs{
        tmp := val
        key := []rune(tmp)
        sort.Slice(key, func(i,j int) bool{
            return byte(key[i]) < byte(key[j])
        })
        m[string(key)] = append(m[string(key)], val)
    }
    ans := make([][]string, 0)
    for _, v := range m{
        ans = append(ans, v)
    }
    return ans
}