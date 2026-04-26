/*
Given an m x n integers matrix, return the length of the longest increasing path in matrix.

From each cell, you can either move in four directions: left, right, up, or down. You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).
*/

func isValid(matrix [][]int, i,j int) bool{
    if i >= 0 && j >=0 && i < len(matrix) && j < len(matrix[0]){
        return true
    }

    return false
}

func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func longestIncreasingPath(matrix [][]int) int {
    dp := make([][]int, len(matrix))
    for i := 0; i < len(matrix); i++{
        dp[i] = make([]int, len(matrix[0]))
    }

    var dfs func(int, int)int
    dfs =  func(i,j int) int{
        if dp[i][j] > 0 {return dp[i][j]}

        best := 1
        dirs := [][2]int{{0,1},{1,0},{0,-1},{-1,0}}
        for _, dir := range dirs{
            ni, nj := i+dir[0], j+dir[1]
            if isValid(matrix, ni,nj) && matrix[ni][nj] > matrix[i][j]{
                best =  max(best, dfs(ni,nj)+1)
            }
        }

        dp[i][j] = best
        return dp[i][j]
    }

    ans := 0
    for i := 0; i < len(matrix); i++{
        for j := 0; j < len(matrix[0]); j++{
            ans = max(ans, dfs(i,j))
        }
    }

    return ans
}