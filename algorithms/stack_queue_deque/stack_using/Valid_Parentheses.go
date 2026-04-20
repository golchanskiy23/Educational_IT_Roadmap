/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

-Open brackets must be closed by the same type of brackets.
-Open brackets must be closed in the correct order.
-Every close bracket has a corresponding open bracket of the same type.
*/

func isValid(s string) bool {
    m := map[rune]rune{
        ')' : '(',
        ']' : '[',
        '}' : '{',
    }

    runes := []rune(s)
    stack := make([]rune, 0)
    for i := 0; i < len(runes); i++{
        c := runes[i]
        if len(stack) == 0{
            stack = append(stack, c)
            continue
        }

        if m[c] == stack[len(stack)-1]{
            stack = stack[:len(stack)-1]
        } else{
            stack = append(stack, c)
        }
    }

    if len(stack) == 0{return true}
    return false
}