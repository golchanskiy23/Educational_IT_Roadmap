/*
В левом верхнем углу прямоугольной таблицы размером 
N
×
M
N×M находится черепашка. В каждой клетке таблицы записано некоторое число. Черепашка может перемещаться вправо или вниз, при этом маршрут черепашки заканчивается в правом нижнем углу таблицы.

Подсчитаем сумму чисел, записанных в клетках, через которую проползла черепашка (включая начальную и конечную клетку). Найдите наибольшее возможное значение этой суммы и маршрут, на котором достигается эта сумма.
*/

package main

import (
  "bufio"
  "os"
  "strconv"
  "strings"
  "fmt"
)

func isValid(i,j,n,m int) bool{
  if i >= 0 && i < n && j < m && j >= 0{
    return true
  }

  return false
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1<<20)
  writer := bufio.NewWriterSize(os.Stdout, 1<<20)
  defer writer.Flush()

  line, _ := reader.ReadString('\n')
  str := strings.Fields(line)
  n, _ := strconv.Atoi(str[0])
  m, _ := strconv.Atoi(str[1])

  type pair struct{
    val, row, col int
  }

  grid, matrix := make([][]pair, n), make([][]int, n)
  for i := 0; i < n; i++{
    line, _ := reader.ReadString('\n')
    str = strings.Fields(line)
    
    grid[i] = make([]pair, m)
    matrix[i] = make([]int, m)
    
    for j := 0; j < m; j++{
      grid[i][j] = pair{}
      val, _ := strconv.Atoi(str[j])
      matrix[i][j] = val
    }
  }

  grid[0][0] = pair{
    val: matrix[0][0],
    row: -1,
    col: -1,
  }
  
  for i := 0; i < n; i++{
    for j := 0; j < m; j++{
      if i == 0 && j == 0{continue}
      if !isValid(i-1,j,n,m){
        grid[i][j] = pair{
          val: grid[i][j-1].val+matrix[i][j],
          row: i,
          col: j-1,
        }
      } else if !isValid(i,j-1,n,m){
        grid[i][j] = pair{
          val: grid[i-1][j].val+matrix[i][j],
          row: i-1,
          col: j,
        } 
      } else{
        curr := pair{}
        if grid[i-1][j].val > grid[i][j-1].val{
          curr.val = grid[i-1][j].val+matrix[i][j]
          curr.row = i-1
          curr.col = j
        } else{
          curr.val = grid[i][j-1].val+matrix[i][j]
          curr.row = i
          curr.col = j-1
        }
        grid[i][j] = curr
      }
    }
  }

  start_i, start_j := n-1, m-1
  stack := []string{}
  for start_i != 0 || start_j != 0{
    if start_i - grid[start_i][start_j].row == 1{
      stack = append(stack, "D")
      start_i = grid[start_i][start_j].row
    } else{
      stack = append(stack, "R")
      start_j = grid[start_i][start_j].col
    }
  }

  fmt.Println(grid[n-1][m-1].val)
  for i := len(stack)-1; i >= 0; i--{
    fmt.Fprint(writer, stack[i], " ")
  }
}