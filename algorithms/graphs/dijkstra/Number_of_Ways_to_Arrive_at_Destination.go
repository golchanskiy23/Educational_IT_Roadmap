/*
You are in a city that consists of n intersections numbered from 0 to n - 1 with bi-directional roads between some intersections. The inputs are generated such that you can reach any intersection from any other intersection and that there is at most one road between any two intersections.

You are given an integer n and a 2D integer array roads where roads[i] = [ui, vi, timei] means that there is a road between intersections ui and vi that takes timei minutes to travel. You want to know in how many ways you can travel from intersection 0 to intersection n - 1 in the shortest amount of time.

Return the number of ways you can arrive at your destination in the shortest amount of time. Since the answer may be large, return it modulo 109 + 7.
*/

type MinHeap [][2]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() any { 
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func countPaths(n int, roads [][]int) int {
    // стоимость пути / количество путей
    dist := make([][2]int, n)
    for i := 1; i < n; i++ {dist[i][0] = math.MaxInt}

    graph := make([][][2]int, n)
    for i := 0; i < len(roads); i++{
        from, to, val := roads[i][0], roads[i][1], roads[i][2]
        graph[from] = append(graph[from], [2]int{to, val})
        graph[to] = append(graph[to], [2]int{from, val})
    }
    dist[0][1] = 1

    // стоимость пути / узел
    h := &MinHeap{{0,0}}
    heap.Init(h)
    for len(*h) > 0{
        front := heap.Pop(h).([2]int)
        cost, node := front[0], front[1]
        if cost > dist[node][0]{
            continue
        }

        for _, neighbor := range graph[node]{
            to, val := neighbor[0], neighbor[1]
            if cost+val < dist[to][0]{
                dist[to][0] = cost+val
                dist[to][1] = dist[node][1]
                heap.Push(h, [2]int{dist[to][0], to})
            } else if cost+val == dist[to][0]{
                dist[to][1] = (dist[to][1]+dist[node][1])% 1000000007
            }
        } 
    }
    
    return dist[n-1][1]
}