/*
Given string num representing a non-negative integer num, and an integer k, return the smallest possible integer after removing k digits from num.
*/

func removeKdigits(num string, k int) string {
    if k >= len(num){
        return "0"
    }
    stack := make([]string, 0)

    for i := 0; i < len(num); i++{
        if len(stack) == 0{
            stack = append(stack, string(num[i]))
            continue
        }

        for len(stack) > 0 && k > 0 && string(num[i]) < stack[len(stack)-1]{
            stack = stack[:len(stack)-1]
            k--
        }

        stack = append(stack, string(num[i]))
    }

    i := 0
    for i < len(stack) && stack[i] == "0"{
        i++
    }
    stack = stack[i:]

    for len(stack) > 0 && k > 0{
        stack = stack[:len(stack)-1]
        k--
    }

    ans := strings.Join(stack, "")
    if ans == ""{return "0"}

    return ans
}