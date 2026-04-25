/*
You are given a network of n nodes, labeled from 1 to n. You are also given times, a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node, vi is the target node, and wi is the time it takes for a signal to travel from source to target.

We will send a signal from a given node k. Return the minimum time it takes for all the n nodes to receive the signal. If it is impossible for all the n nodes to receive the signal, return -1.
*/

type pair struct{
    // номер узла(конечного), расстояние от k
    to, val int
}

type MinHeap []pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(pair))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func max(a,b int) int{
    if a > b{
        return a
    }

    return b
}

// алгоритм Дейкстры - находим кратчайший путь от одной до всех остальных вершин
func networkDelayTime(times [][]int, n int, k int) int {
    // строим adjacent_list
    adj := make([][]pair, n+1)
    for i := 1; i <= n; i++{
        adj[i] = make([]pair, 0)
    }

    for i := 0; i < len(times); i++{
        adj[times[i][0]] = append(adj[times[i][0]], pair{
        to:  times[i][1],
        val: times[i][2],
        })
    }

    dist := make([]int, n+1)
    for i := 1; i <= n; i++ {dist[i] = math.MaxInt}
    dist[k] = 0
    // закидывыаем в min_heap(хранит расстояние и номер узла) первую
    minHeap := &MinHeap{}
    heap.Push(minHeap, pair{
        to: k,
        val: 0,
    })

    // пока очередь не пуста, берём минимальную по весу ноду, считаю newDist = dist[curr] + weight и смотрю меньше ли newDist чем dist[to]
    for len(*minHeap) > 0{
        front := heap.Pop(minHeap).(pair)
        if front.val > dist[front.to]{
            continue
        }

        for _, edge := range adj[front.to]{
            newDist := dist[front.to]+edge.val
            // и если да, обновляем и закидываем в heap
            if newDist < dist[edge.to]{
                dist[edge.to] = newDist
                heap.Push(minHeap, pair{
                    to: edge.to,
                    val: newDist,
                })
            }
        }
    }

    ans := 0
    for i := 1; i <= n; i++{
        if dist[i] == math.MaxInt { return -1 }
        ans = max(ans, dist[i])
    }

    return ans
}