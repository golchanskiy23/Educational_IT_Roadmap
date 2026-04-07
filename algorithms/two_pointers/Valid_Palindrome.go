/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.
*/

func isPalindrome(s string) bool {
    str := strings.ToLower(s)
    arr := []rune(str)
    app := make([]rune, 0)
    
    for i := 0; i < len(str); i++{
        if (byte(arr[i]) >= 97 && byte(arr[i]) <= 122) ||
        (byte(arr[i]) >= 48 && byte(arr[i]) <= 57){
            app = append(app, arr[i])
        }
    }
   // fmt.Println(arr)
    left,right := 0, len(app)-1
    for left < right{
        if app[left] != app[right]{
            return false
        }
        left++
        right--
    }
    return true
}