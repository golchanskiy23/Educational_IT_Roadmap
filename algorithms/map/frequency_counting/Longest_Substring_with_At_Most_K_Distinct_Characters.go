/*
You are given a string s and an integer k. Your task is to find the length of the longest substring within s that contains at most k distinct characters.

A substring is a contiguous sequence of characters within the string. For example, if s = "eceba" and k = 2, you need to find the longest substring that has no more than 2 different characters.
*/

func max(a,b int) int{
	if a > b{
		return a
	}
	return b
}

func lengthOfLongestSubstringKDistinct(s string, k int) int{
	m := make(map[int]int)
	runes := []rune(s)
	l, r := 0,0
	ans := 0
	for r < len(s){
		m[runes[r]]++
		if len(m) <= k{
			ans = max(ans, r-l+1)
		}

		for len(m) > k{
			m[runes[l]]--
			if m[runes[l]] == 0{
				delete(m, runes[l])
			}
			l++
		}
	}

	return ans
}