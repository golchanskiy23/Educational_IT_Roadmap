/*
You are a hiker preparing for an upcoming hike. You are given heights, a 2D array of size rows x columns, where heights[row][col] represents the height of cell (row, col). You are situated in the top-left cell, (0, 0), and you hope to travel to the bottom-right cell, (rows-1, columns-1) (i.e., 0-indexed). You can move up, down, left, or right, and you wish to find a route that requires the minimum effort.

A route's effort is the maximum absolute difference in heights between two consecutive cells of the route.

Return the minimum effort required to travel from the top-left cell to the bottom-right cell.
*/

type triplet struct{
    r, c, effort int
}

type MinHeap []triplet

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].effort < h[j].effort }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(triplet))
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

func abs(a int) int{
    if a < 0{return (-1)*a}
    return a
}

func isValid(grid [][]int, i,j int)bool{
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]){
        return true
    }

    return false
}

func minimumEffortPath(heights [][]int) int {
    dist := make([][]int, len(heights))
    for i := 0; i < len(heights); i++{
        dist[i] = make([]int, len(heights[0]))
    }

    for i := 0; i < len(heights); i++{
        for j := 0; j < len(heights[0]); j++{
            dist[i][j] = math.MaxInt
        }
    }

    dist[0][0] = 0
    dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}
    minHeap := &MinHeap{}
    heap.Push(minHeap, triplet{
        r: 0,
        c: 0,
        effort: 0,
    })

    for len(*minHeap) > 0{
        front := heap.Pop(minHeap).(triplet)
        if front.effort > dist[front.r][front.c]{
            continue
        }

        for _, dir := range dirs{
            if !isValid(heights, front.r+dir[0], front.c+dir[1]){
                continue
            }
            newr, newc := front.r+dir[0], front.c+dir[1]
            newMax := max(front.effort, abs(heights[newr][newc]-heights[front.r][front.c]))
            if newMax < dist[newr][newc]{
                dist[newr][newc] = newMax
                heap.Push(minHeap, triplet{
                    r: newr,
                    c: newc,
                    effort: newMax,
                })
            }
        }
    }

    return dist[len(heights)-1][len(heights[0])-1]
}