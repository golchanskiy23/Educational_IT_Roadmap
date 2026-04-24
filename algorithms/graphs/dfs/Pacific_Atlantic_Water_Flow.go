/*
There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.
*/

type ceil struct{
        i,j int
}

func isValid(grid [][]int, i,j int)bool{
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]){
        return true
    }

    return false
}

func dfs(grid [][]int, i, j, prevHeight int, seen, set map[ceil]bool) {
    curr := ceil{i, j}
    if !isValid(grid, i, j) || grid[i][j] < prevHeight || seen[curr] {
        return
    }
    seen[curr] = true
    set[curr] = true
    dirs := [][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
    for _, dir := range dirs {
        dfs(grid, i+dir[0], j+dir[1], grid[i][j], seen, set)
    }
}

func pacificAtlantic(grid [][]int) [][]int {
    n, m := len(grid), len(grid[0])
    pacific_set, atlantic_set := make(map[ceil]bool), make(map[ceil]bool)
    pacificSeen, atlanticSeen := make(map[ceil]bool), make(map[ceil]bool)
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if i == 0 || j == 0{
                dfs(grid, i,j, 0, pacificSeen, pacific_set)
            }

            if i == n-1 || j == m-1{
                dfs(grid, i,j, 0, atlanticSeen, atlantic_set)
            }
        }
    }

    var ans [][]int
    for k,_ := range pacific_set{
        if _, ok := atlantic_set[k]; ok{
            tmp := []int{k.i, k.j}
            ans  = append(ans, tmp)
        }
    }

    return ans
}