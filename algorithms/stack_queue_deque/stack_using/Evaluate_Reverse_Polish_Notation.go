/*
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

-The valid operators are '+', '-', '*', and '/'.
-Each operand may be an integer or another expression.
-The division between two integers always truncates toward zero.
-There will not be any division by zero.
-The input represents a valid arithmetic expression in a reverse polish notation.
-The answer and all the intermediate calculations can be represented in a 32-bit integer.
*/

func calc(f,s int, t string) int{
    ans := 0

    switch t{
        case "-":
            ans = f-s
        case "+":
            ans = f+s
        case "/":
            ans = f/s
        case "*":
            ans = f*s
    }

    return ans
}

func evalRPN(tokens []string) int {
    ans := 0
    stack := make([]int, 0)

    for i := 0; i < len(tokens); i++{
        num, err := strconv.Atoi(tokens[i])
        if err != nil{
            s := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            f := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans = calc(f,s,tokens[i])
            stack = append(stack, ans)
        } else{
            stack = append(stack, num)
        }
    }

    return stack[0]
}