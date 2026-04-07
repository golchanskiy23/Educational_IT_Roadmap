/*
You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.
Return the max sliding window.
*/

import "container/heap"

type MaxHeap [][2]int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.([2]int))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func maxSlidingWindow(nums []int, k int) []int {
    l,r,n := 0, k-1, len(nums)
    if n == 1 {return []int{nums[0]}}

    ans := make([]int, 0 ,n-k+1)
    h := &MaxHeap{}
    heap.Init(h)
    for i := l; i <= r; i++{
        heap.Push(h, [2]int{nums[i], i})
    }

    for r < n{
        for (*h)[0][1] < l{
            heap.Pop(h)
        }
        ans = append(ans, (*h)[0][0])
        l++
        r++
        if r == n{break}
        heap.Push(h, [2]int{nums[r],r})
    }

    return ans
}