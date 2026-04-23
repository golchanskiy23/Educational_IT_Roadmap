/*
The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.
*/

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type MedianFinder struct {
    minHeap *MinHeap
    maxHeap *MaxHeap
}


func Constructor() MedianFinder {
    return MedianFinder{
        minHeap: &MinHeap{},
        maxHeap: &MaxHeap{},
    }    
}


func (this *MedianFinder) AddNum(num int)  {
    if len(*this.maxHeap) == 0{
        heap.Push(this.maxHeap, num)
        return
    }

    max_top := (*this.maxHeap)[0]
    if num <= max_top{
        heap.Push(this.maxHeap, num)
    } else{
        heap.Push(this.minHeap, num)
    }
    if len(*this.minHeap) > len(*this.maxHeap){
        heap.Push(this.maxHeap, heap.Pop(this.minHeap).(int))
    } else if len(*this.maxHeap)-len(*this.minHeap) > 1 {
        heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    max := (*this.maxHeap)[0]
    if len(*this.minHeap) == len(*this.maxHeap){
        min := (*this.minHeap)[0]
        return (float64(min)+float64(max))/2.0
    }

    return float64(max)
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */