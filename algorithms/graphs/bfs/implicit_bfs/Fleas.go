package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(p struct{ x, y int }, n, m int) bool {
	return p.x >= 1 && p.x <= n && p.y >= 1 && p.y <= m
}

func main() {
	type point struct {
		x, y int
	}

	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.TrimSpace(firstLine)
	parts := strings.Fields(firstLine)

	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	s, _ := strconv.Atoi(parts[2])
	t, _ := strconv.Atoi(parts[3])
	q, _ := strconv.Atoi(parts[4])

    set := make(map[point]struct{})
	for i := 0; i < q; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		coords := strings.Fields(line)
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		set[point{x, y}] = struct{}{}
	}

	eating := point{s, t}
	dist := make(map[point]int)
    dist[eating] = 0

	queue := []point{eating}
	dirs := [][2]int{{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				cx, cy := top.x+dir[0], top.y+dir[1]
				p := point{cx, cy}
				if isValid(p, n, m) {
					if _, seen := dist[p]; !seen{
                        dist[p] = dist[top]+1
                        queue = append(queue, p)
                    }
				}
			}
		}
	}

	cnt := 0
    for f, _ := range set{
        total, ok := dist[f]
		if !ok{
            fmt.Println(-1)
            return
        }
        cnt += total
    }

	fmt.Println(cnt)
}