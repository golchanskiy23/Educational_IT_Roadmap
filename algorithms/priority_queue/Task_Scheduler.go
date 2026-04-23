/*
You are given an array of CPU tasks, each labeled with a letter from A to Z, and a number n. Each CPU interval can be idle or allow the completion of one task. Tasks can be completed in any order, but there's a constraint: there has to be a gap of at least n intervals between two tasks with the same label.

Return the minimum number of CPU intervals required to complete all tasks.
*/

type Pair struct{
    letter byte
    amount int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].amount > h[j].amount }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func leastInterval(tasks []byte, n int) int {
    m := make(map[byte]int)
    for i := 0;i < len(tasks); i++{
        m[tasks[i]]++
    }
    h := &MaxHeap{}
    for k,v := range m{
        heap.Push(h, Pair{
            letter: k,
            amount: v,
        })       
    }

    var ans int
    for len(*h) > 0{
        cycle := min(n+1, len(*h))
        tmp := make([]Pair, 0)
        ans += (n+1)
        for i := 0; i < cycle; i++{
            pair := heap.Pop(h).(Pair)
            pair.amount--
            if pair.amount > 0{
                tmp = append(tmp, pair)   
            }
        }
        if len(tmp) == 0 && len(*h) == 0{
            ans -= (n+1-cycle)
        }

        for _, val := range tmp{
            heap.Push(h, val)
        }
    }

    return ans
}