/*
You are given an n x n binary matrix grid. You are allowed to change at most one 0 to be 1.

Return the size of the largest island in grid after applying this operation.

An island is a 4-directionally connected group of 1s.
*/

func max(a,b int) int{
    if a > b{
        return a
    }

    return b
}

func isValid(grid [][]int, i,j int)bool{
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]){
        return true
    }

    return false
}

func dfs(grid [][]int, i,j,idx int) int{
    if !isValid(grid, i,j) || grid[i][j] != 1 {
        return 0
    }

    grid[i][j] = idx

    size := 1

    dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}
    for _, d := range dirs {
        size += dfs(grid, i+d[0], j+d[1], idx)
    }

    return size
}

func largestIsland(grid [][]int) int {
    // сначала dfs по всем единицам и увеличиваем индекс и закидываем в map от idx
    n,m, idx := len(grid), len(grid[0]), 2
    mi := make(map[int]int)
    for i := 0; i < n ; i++{
        for j := 0; j < m; j++{
            if grid[i][j] == 1{
                size := dfs(grid, i, j, idx)
                mi[idx] = size
                idx++
            }
        }
    }
    ans := 0    
    for _, v := range mi {
        ans = max(ans, v)
    }

    // затем проходимся по всем нулям и смотрим на соседние узлы и образуем общий остров
    dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if grid[i][j] == 0{
                curr := 1
                seen := make(map[int]bool)
                for _, d := range dirs {
                    if isValid(grid, i+d[0], j+d[1]){
                        id := grid[i+d[0]][j+d[1]]
                        if !seen[id] && id > 1{
                            seen[id] = true
                            curr += mi[id]
                        }
                    }      
                }
                ans = max(ans, curr)
            }
        }
    }
    
    return ans
}