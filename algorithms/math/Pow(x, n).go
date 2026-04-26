/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/

func util(x float64, n int) float64{
    tmp := 1.0
    for n != 0{
        if n % 2 != 0{
           tmp *= x
        }
        x *= x
        n /= 2
    }

    return tmp
}

func myPow(x float64, n int) float64 {
    if n == 0{return 1}
    isNegative := false
    if n < 0{
        isNegative = true
        n *= (-1)
    }
    t := util(x,n)
    
    if isNegative{
        return 1/t
    }

    return t
}