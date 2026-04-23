/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.
*/

func minEatingSpeed(piles []int, h int) int {
    sort.Ints(piles)
    l, r := 1, 1000000000
    for l < r{
        mid := l + (r-l)/2
        sum := 0
        for i := 0; i < len(piles); i++{
            rem := piles[i]%mid
            sum += (piles[i]/mid)
            if rem != 0 {sum += 1}
        }
        if sum <= h{
            r = mid
        } else{
            l = mid+1
        }
    }

    return l
}