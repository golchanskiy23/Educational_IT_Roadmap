/*
В каждой клетке прямоугольной таблицы 
N
×
M
N×M записано некоторое число. Изначально игрок находится в левой верхней клетке. За один ход ему разрешается перемещаться в соседнюю клетку либо вправо, либо вниз (влево и вверх перемещаться запрещено). При проходе через клетку с игрока берут столько килограммов еды, какое число записано в этой клетке (еду берут также за первую и последнюю клетки его пути).

Требуется найти минимальный вес еды в килограммах, отдав которую игрок может попасть в правый нижний угол.
*/

package main

import (
  "bufio"
  "os"
  "strconv"
  "strings"
  "math"
)

func isValid(i,j,n,m int) bool{
  if i >= 0 && i < n && j < m && j >= 0{
    return true
  }

  return false
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1<<20)
  writer := bufio.NewWriterSize(os.Stdout, 1<<20)
  defer writer.Flush()

  line, _ := reader.ReadString('\n')
  str := strings.Fields(line)
  n, _ := strconv.Atoi(str[0])
  m, _ := strconv.Atoi(str[1])

  grid, matrix := make([][]int, n), make([][]int, n)
  for i := 0; i < n; i++{
    line, _ := reader.ReadString('\n')
    str = strings.Fields(line)
    
    grid[i] = make([]int, m)
    matrix[i] = make([]int, m)
    
    for j := 0; j < m; j++{
      grid[i][j] = math.MaxInt
      val, _ := strconv.Atoi(str[j])
      matrix[i][j] = val
    }
  }

  grid[0][0] = matrix[0][0]
  for i := 0; i < n; i++{
    for j := 0; j < m; j++{
      curr := grid[i][j]
      up, left := math.MaxInt, math.MaxInt
      if isValid(i-1, j, n,m){
        up = grid[i-1][j]+matrix[i][j]
        curr = min(curr, up)
      }

      if isValid(i, j-1, n,m){
        left = grid[i][j-1]+matrix[i][j]
        curr = min(curr, left)
      }
      grid[i][j] = curr
    }
  }

  writer.WriteString(strconv.Itoa(grid[n-1][m-1]))
  writer.WriteByte('\n')
}