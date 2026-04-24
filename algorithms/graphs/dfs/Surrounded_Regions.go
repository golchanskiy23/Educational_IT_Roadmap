/*
You are given an m x n matrix board containing letters 'X' and 'O', capture regions that are surrounded:

Connect: A cell is connected to adjacent cells horizontally or vertically.
Region: To form a region connect every 'O' cell.
Surround: A region is surrounded if none of the 'O' cells in that region are on the edge of the board. Such regions are completely enclosed by 'X' cells.
To capture a surrounded region, replace all 'O's with 'X's in-place within the original board. You do not need to return anything.
*/

func isValid(grid [][]byte, i,j int)bool{
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]){
        return true
    }

    return false
}

func isEdge(board [][]byte, i,j int) bool{
    if i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1{
        return true
    }

    return false
}

func dfs(grid *[][]byte, i,j int){
    if !isValid(*grid, i,j) || (*grid)[i][j] != 'Y'{
        return
    }

    (*grid)[i][j] = 'O'
    dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}
    for _, d := range dirs {
        dfs(grid, i+d[0], j+d[1])
    }
}

// могут быть 'O' на границе
// ['X' 'X' 'X' 'X']
// ['X' 'O' 'X' 'X']
// ['O' 'O' 'X' 'X]
func solve(board [][]byte)  {
    // меняем все 'O' на 'Y'
    n,m := len(board), len(board[0])
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if board[i][j] == 'O'{board[i][j] = 'Y'}
        }
    }

    // проходимся по матрице, запускаемся от границ , начинающихся с 'Y' и меняем их на 'O'
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if board[i][j] == 'X'{continue}
            if board[i][j] == 'Y' && isEdge(board, i,j){
                dfs(&board, i,j)
            }
        }
    }

    // меняем все 'Y' на 'X'
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if board[i][j] == 'Y'{
                board[i][j] = 'X'
            }
        }
    }
}