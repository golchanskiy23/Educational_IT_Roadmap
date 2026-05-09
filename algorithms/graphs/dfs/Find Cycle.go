/*
Дан неориентированный граф. Требуется определить, есть ли в нем цикл, и, если есть, вывести его.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n          int
	adj        [][]int
	color      []int
	parent     []int
	cycleStart int
	cycleEnd   int
)

func dfs(v, par int) {
	if cycleStart != -1 {
		return
	}
	color[v] = 1
	for _, u := range adj[v] {
		if u == par {
			continue
		}
		if color[u] == 1 {
			cycleStart = u
			cycleEnd = v
			return
		}
		if color[u] == 0 {
			parent[u] = v
			dfs(u, v)
		}
		if cycleStart != -1 {
			return
		}
	}
	color[v] = 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &n)

	adj = make([][]int, n+1)
	color = make([]int, n+1)
	parent = make([]int, n+1)
	cycleStart = -1
	cycleEnd = -1

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var val int
			fmt.Fscan(reader, &val)
			if val == 1 && j > i {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	for s := 1; s <= n; s++ {
		if color[s] == 0 && cycleStart == -1 {
			dfs(s, -1)
		}
		if cycleStart != -1 {
			break
		}
	}

	if cycleStart == -1 {
		fmt.Fprintln(writer, "NO")
		return
	}

	path := []int{}
	for v := cycleEnd; v != cycleStart; v = parent[v] {
		path = append(path, v)
	}
	path = append(path, cycleStart)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	fmt.Fprintln(writer, "YES")
	fmt.Fprintln(writer, len(path))
	for i, v := range path {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}