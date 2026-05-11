package main

import (
  "bufio"
  "os"
  "strconv"
  "strings"
  "fmt"
)

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1<<20)
  writer := bufio.NewWriterSize(os.Stdout, 1<<20)
  defer writer.Flush()

  line, _ := reader.ReadString('\n')
  n, _ := strconv.Atoi(strings.TrimSpace(line))

  graph := make([][]int, n+1)
  for i := 0; i <= n; i++{
    graph[i] = make([]int, n+1)
  }

  for i := 1; i <= n; i++{
    for j := 1; j <= n; j++{
      var val int
			fmt.Fscan(reader, &val)
			if val == 1 && j > i {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
    }
  }

  var val int
	fmt.Fscan(reader, &val)
  a := val
  fmt.Fscan(reader, &val)
  b := val


  if a == b{
    writer.WriteString(strconv.Itoa(0))
    writer.WriteByte('\n')
    return
  }

  parent, visited, dist := make([]int, n+1), make([]bool, n+1), make([]int, n+1)
  for i := range dist {
    dist[i] = -1
  }

  dist[a] = 0
  parent[a] = -1
  visited[a] = true
  dist[a] = 0
  
  queue := []int{a}
  for len(queue) > 0{
    size := len(queue)
    for i := 0; i < size; i++{
      top := queue[0]
      queue = queue[1:]
      for _, neighbor := range graph[top]{
        if visited[neighbor]{
          continue
        }

        visited[neighbor] = true
        parent[neighbor] = top
        dist[neighbor] = dist[top]+1
        queue = append(queue, neighbor)
      }
    }
  }

  if dist[b] == -1{
    writer.WriteString(strconv.Itoa(-1))
    writer.WriteByte('\n')
    return
  }

  ans := make([]int, 0, dist[b]+1)
  for v := b; v != a; v = parent[v]{
    ans = append(ans, v)
  }

  ans = append(ans, a)

  fmt.Fprintln(writer, len(ans)-1)

  for i := len(ans)-1; i >= 0; i--{
    fmt.Fprint(writer, ans[i])
    if i > 0 {
			fmt.Fprint(writer, " ")
		}
  }
}