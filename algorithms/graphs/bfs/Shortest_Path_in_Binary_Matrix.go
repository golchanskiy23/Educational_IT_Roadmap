/*
Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If there is no clear path, return -1.

A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:

All the visited cells of the path are 0.
All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).
The length of a clear path is the number of visited cells of this path.
*/

func isValid(grid [][]int, i,j int)bool{
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]){
        return true
    }

    return false
}

func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] == 1{return -1}
    ans := make([][]int, len(grid))
    queue := [][2]int{}
    for i := 0; i < len(grid); i++{
        ans[i] = make([]int, len(grid[0]))
        for j := 0; j < len(grid[0]); j++{
            ans[i][j] = -1
        }
    }

    ans[0][0] = 1
    queue = append(queue, [2]int{0, 0}) 

    for len(queue) > 0{
        size := len(queue)
        for i := 0; i < size; i++{
            front := queue[0]
            queue = queue[1:]
            idx_i, idx_j := front[0], front[1]

            if isValid(grid, idx_i, idx_j+1) && grid[idx_i][idx_j+1] == 0 && ans[idx_i][idx_j+1] == -1{
                ans[idx_i][idx_j+1] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i, idx_j+1})
            }

            if isValid(grid, idx_i-1, idx_j+1) && grid[idx_i-1][idx_j+1] == 0 && ans[idx_i-1][idx_j+1] == -1{
                ans[idx_i-1][idx_j+1] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i-1, idx_j+1})
            }

            if isValid(grid, idx_i-1, idx_j) && grid[idx_i-1][idx_j] == 0 && ans[idx_i-1][idx_j] == -1{
                ans[idx_i-1][idx_j] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i-1, idx_j})
            }

            if isValid(grid, idx_i-1, idx_j-1) && grid[idx_i-1][idx_j-1] == 0 && ans[idx_i-1][idx_j-1] == -1{
                ans[idx_i-1][idx_j-1] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i-1, idx_j-1})
            }

            if isValid(grid, idx_i, idx_j-1) && grid[idx_i][idx_j-1] == 0 && ans[idx_i][idx_j-1] == -1{
                ans[idx_i][idx_j-1] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i, idx_j-1})
            }

            if isValid(grid, idx_i+1, idx_j-1) && grid[idx_i+1][idx_j-1] == 0 && ans[idx_i+1][idx_j-1] == -1{
                ans[idx_i+1][idx_j-1] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i+1, idx_j-1})
            }

            if isValid(grid, idx_i+1, idx_j) && grid[idx_i+1][idx_j] == 0 && ans[idx_i+1][idx_j] == -1{
                ans[idx_i+1][idx_j] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i+1, idx_j})
            }

            if isValid(grid, idx_i+1, idx_j+1) && grid[idx_i+1][idx_j+1] == 0 && ans[idx_i+1][idx_j+1] == -1{
                ans[idx_i+1][idx_j+1] = ans[idx_i][idx_j]+1
                queue = append(queue, [2]int{idx_i+1, idx_j+1})
            }
        }
    }

    return ans[len(grid)-1][len(grid[0])-1]
}