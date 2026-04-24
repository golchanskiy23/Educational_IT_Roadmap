/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.
*/

func isValid(grid [][]int, i,j int)bool{
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]){
        return true
    }

    return false
}

func orangesRotting(grid [][]int) int {
    queue := [][2]int{}
    for i, row := range grid{
        for j, val := range row{
            if val == 2{
                queue = append(queue, [2]int{i,j})
            }
        } 
    }

    ans := 0
    for len(queue) > 0{
        size := len(queue)
        for i := 0; i < size; i++{
            front := queue[0]
            queue = queue[1:]
            idx_i, idx_j := front[0], front[1]

            if isValid(grid, idx_i+1, idx_j) && grid[idx_i+1][idx_j] == 1 {
                queue = append(queue, [2]int{idx_i+1, idx_j})
                grid[idx_i+1][idx_j] = 2   
            }
            if isValid(grid, idx_i-1, idx_j) && grid[idx_i-1][idx_j] == 1 {
                queue = append(queue, [2]int{idx_i-1, idx_j})
                grid[idx_i-1][idx_j] = 2
            }
            if isValid(grid, idx_i, idx_j-1) && grid[idx_i][idx_j-1] == 1 {
                queue = append(queue, [2]int{idx_i, idx_j-1})
                grid[idx_i][idx_j-1] = 2
            }
            if isValid(grid, idx_i, idx_j+1) && grid[idx_i][idx_j+1] == 1 {
                queue = append(queue, [2]int{idx_i, idx_j+1})
                grid[idx_i][idx_j+1] = 2
            }
        }
        if len(queue) > 0 {ans++}
    }

    for _, row := range grid{
        for _, val := range row{
            if val == 1{
                return -1
            }
        }
    }

    return ans
}