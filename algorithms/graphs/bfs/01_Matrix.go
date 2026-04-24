/*
Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.

The distance between two cells sharing a common edge is 1.
*/

func isValid(grid [][]int, i,j int)bool{
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]){
        return true
    }

    return false
}

func updateMatrix(grid [][]int) [][]int {
    ans := make([][]int, len(grid))
    queue := [][2]int{}
    for i := 0; i < len(grid); i++{
        ans[i] = make([]int, len(grid[0]))
        for j := 0; j < len(grid[0]); j++{
            ans[i][j] = -1
            if grid[i][j] == 0{
                queue = append(queue, [2]int{i,j})
                ans[i][j] = 0
            }
        }
    }

    for len(queue) > 0{
        size := len(queue)
        for i := 0; i < size; i++{
            front := queue[0]
            queue = queue[1:]
            idx_i, idx_j := front[0], front[1]

            if isValid(grid, idx_i+1, idx_j) && ans[idx_i+1][idx_j] == -1{
                ans[idx_i+1][idx_j] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i+1, idx_j})
            }

            if isValid(grid, idx_i-1, idx_j) && ans[idx_i-1][idx_j] == -1{
                ans[idx_i-1][idx_j] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i-1, idx_j})
            }

            if isValid(grid, idx_i, idx_j+1) && ans[idx_i][idx_j+1] == -1{
                ans[idx_i][idx_j+1] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i, idx_j+1})
            }

            if isValid(grid, idx_i, idx_j-1) && ans[idx_i][idx_j-1] == -1{
                ans[idx_i][idx_j-1] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i, idx_j-1})
            }
        }
    }

    return ans
}