/*
We are given an array asteroids of integers representing asteroids in a row. The indices of the asteroid in the array represent their relative position in space.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.
*/

func isSameSign(a,b int) bool{
    if (a > 0 && b > 0) || (a < 0 && b < 0){
        return true
    }
    return false
}

func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0)
    for i := 0; i < len(asteroids); i++{
        if len(stack) == 0{
            stack = append(stack, asteroids[i])
            continue
        }

        if asteroids[i] > 0 || isSameSign(asteroids[i], stack[len(stack)-1]){
            stack = append(stack, asteroids[i])
            continue
        }

        flag := true
        for len(stack) > 0{
            if asteroids[i] + stack[len(stack)-1] == 0{
                stack = stack[:len(stack)-1]
                flag = false
                break
            } else if asteroids[i] + stack[len(stack)-1] > 0{
                flag = false
                break
            } else{
                if !isSameSign(asteroids[i], stack[len(stack)-1]){
                    stack = stack[:len(stack)-1]
                } else{
                    break
                }
            }
        }
        if flag{stack = append(stack, asteroids[i])}
    }

    return stack
}