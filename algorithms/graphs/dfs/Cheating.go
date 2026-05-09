/*
Во время контрольной работы профессор Флойд заметил, что некоторые студенты обмениваются записками. Сначала он хотел поставить им всем двойки, но в тот день профессор был добрым, а потому решил разделить студентов на две группы: списывающих и дающих списывать, и поставить двойки только первым.

У профессора записаны все пары студентов, обменявшихся записками. Требуется определить, сможет ли он разделить студентов на две группы так, чтобы любой обмен записками осуществлялся от студента одной группы студенту другой группы.
*/

package main

import (
  "bufio"
  "os"
  "strconv"
  "strings"
)

func dfs(adj [][]int, deg []int, node, val int) bool{
  if deg[node] != -1{
    if deg[node] == val{
      return true
    }
    return false
  }

  deg[node] = val

  for _, neighbor := range adj[node]{
    v := (val+1)%2
    if !dfs(adj, deg, neighbor, v){
      return false
    }
  }

  return true
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1<<20)
  writer := bufio.NewWriterSize(os.Stdout, 1<<20)
  defer writer.Flush()

  line, _ := reader.ReadString('\n')
  arr := strings.Split(line, " ")
  n, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
  m, _ := strconv.Atoi(strings.TrimSpace(arr[1]))

  adj, degree := make([][]int, n+1), make([]int, n+1)
  for i := 1; i <= n; i++{
    adj[i] = make([]int, 0)
    degree[i] = -1
  }

  for i := 0; i < m; i++{
    line, _ = reader.ReadString('\n')
    arr = strings.Split(line, " ")
    v1, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
    v2, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
    adj[v1] = append(adj[v1], v2)
    adj[v2] = append(adj[v2], v1)
  }

  for i := 1; i <= n; i++{
    if degree[i] == -1{
      if !dfs(adj, degree, i, 0){
        writer.WriteString("NO")
        writer.WriteByte('\n')
        return
      }
    }
  }

  writer.WriteString("YES")
  writer.WriteByte('\n')
}