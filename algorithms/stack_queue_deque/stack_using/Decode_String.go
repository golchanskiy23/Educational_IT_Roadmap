/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.
*/

func decodeString(s string) string {
    str, _ := decode(s,0)
    return str
}

func decode(s string, i int) (string, int){
    ans, num := "", 0
    k := i
    for k < len(s){
        if isdigit(s[k]){
            num = 0
            for k < len(s) && isdigit(s[k]){
                num = (num*10)+int(s[k]-'0')
                k++
            }
            continue
        } else if s[k] == '['{
            sub, j := decode(s, k+1)
            ans += repeat(sub, num)
            num = 0
            k = j+1
            continue
        } else if s[k] == ']'{
            return ans, k
        } else{
            ans += string(s[k])
            k++
        }
    }
    return ans, len(ans)
}

func isdigit(num byte) bool{
    c := int(num-'0')
    if c >= 0 && c <= 9{
        return true
    }
    return false
}

func repeat(s string, num int) string{
    res := ""
	for i := 0; i < num; i++ {
		res += s
	}
	return res
}