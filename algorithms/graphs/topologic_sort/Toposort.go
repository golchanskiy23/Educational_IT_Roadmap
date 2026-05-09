/*
Дан ориентированный граф. Необходимо построить топологическую сортировку.

Напомним:

Топологическая сортировка указывает такой линейный порядок на его вершинах, что любое ребро ведёт от вершины с меньшим номером к вершине с большим номером.
*/

package main

import (
  "bufio"
  "os"
  "strconv"
  "strings"
)

var(
  hasCycle bool
)

func dfs(adj [][]int, stack *[]int, color []int, node int){
  if hasCycle{
    return
  }

  color[node] = 1

  for _, neighbor := range adj[node]{
    if color[neighbor] == 1{
      hasCycle = true
      return
    }
    if color[neighbor] == 0{
      dfs(adj, stack, color, neighbor)
    }
  }

  color[node] = 2
  *stack = append(*stack, node)
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1<<20)
  writer := bufio.NewWriterSize(os.Stdout, 1<<20)
  defer writer.Flush()

  line, _ := reader.ReadString('\n')
  arr := strings.Fields(line)
  n, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
  m, _ := strconv.Atoi(strings.TrimSpace(arr[1]))

  adj := make([][]int, n+1)
  color := make([]int, n+1)
  for i := 1; i <= n; i++{
    adj[i] = make([]int, 0)
  }

  for i := 0; i < m; i++{
    line, _ = reader.ReadString('\n')
    arr = strings.Fields(line)
    v1, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
    v2, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
    adj[v1] = append(adj[v1], v2)
  }

  stack := []int{}
  for i := 1; i <= n; i++{
    if color[i] == 0{
      dfs(adj, &stack, color, i)
    }
    if hasCycle{
        writer.WriteString(strconv.Itoa(-1))
        writer.WriteByte('\n')
        return
    }
  }

  for i := len(stack)-1; i >= 0; i--{
    writer.WriteString(strconv.Itoa(stack[i]))
    if i != 0{writer.WriteString(" ")}
  }
  writer.WriteByte('\n')
}