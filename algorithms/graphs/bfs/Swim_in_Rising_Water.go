/*
You are given an n x n integer matrix grid where each value grid[i][j] represents the elevation at that point (i, j).

It starts raining, and water gradually rises over time. At time t, the water level is t, meaning any cell with elevation less than equal to t is submerged or reachable.

You can swim from a square to another 4-directionally adjacent square if and only if the elevation of both squares individually are at most t. You can swim infinite distances in zero time. Of course, you must stay within the boundaries of the grid during your swim.

Return the minimum time until you can reach the bottom right square (n - 1, n - 1) if you start at the top left square (0, 0).
*/

func isValid(grid [][]int, i,j int)bool{
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]){
        return true
    }

    return false
}

func canReach(grid [][]int, n, mid int) bool{
    if grid[0][0] > mid{return false}
    dist := make([][]bool, n)
    for i := 0; i < n; i++ {dist[i] = make([]bool, n)}
    type pair struct{
        i,j int
    }

    queue := []pair{pair{
        i: 0,
        j: 0,
    }}
    dist[0][0] = true
    for len(queue) > 0{
        front := queue[0]
        queue = queue[1:]
        dirs := [][2]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
        for _, dir := range dirs{
            ni, nj := front.i+dir[0], front.j+dir[1]
            if isValid(grid, ni, nj) && grid[ni][nj] <= mid && !dist[ni][nj]{
                dist[ni][nj] = true
                queue = append(queue, pair{
                    i: ni,
                    j: nj,
                })
            }
        }
    }
    fmt.Println(mid)
    for i := 0; i < n; i++{
        for j := 0; j < n; j++{
            fmt.Print(dist[i][j], " ")
        }
        fmt.Println()
    }

    return dist[n-1][n-1]
}

func swimInWater(grid [][]int) int {
    n := len(grid)

    l,r := 0, n*n-1
    for l < r{
        mid := l + (r-l)/2
        if canReach(grid, n, mid){
            r = mid
        } else{
            l = mid+1
        }
        fmt.Println()
    }


    return l
}