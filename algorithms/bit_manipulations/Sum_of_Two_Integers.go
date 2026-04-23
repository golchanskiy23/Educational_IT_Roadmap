/*
Given two integers a and b, return the sum of the two integers without using the operators + and -.
*/

func getSum(a int, b int) int {
    x,y := uint32(a), uint32(b)
    for y != 0{
        carry := x & y
        x = x ^ y
        y = carry << 1
    }
    return int(int32(x))
}