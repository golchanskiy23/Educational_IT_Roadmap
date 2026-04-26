/*
Given an integer n, return the number of prime numbers that are strictly less than n.
*/

func isPrime(n int, primes []int) bool{
    for _, p := range primes {
        if p*p > n {
            break
        }
        if n%p == 0 {
            return false
        }
    }
    return true
}

func countPrimes(n int) int {
    if n == 0 || n == 1 || n == 2 {return 0}
    primes := []int{2}
    for i := 3; i < n; i++{
        if isPrime(i, primes) {
            primes = append(primes, i)
        }
    }

    return len(primes)
}