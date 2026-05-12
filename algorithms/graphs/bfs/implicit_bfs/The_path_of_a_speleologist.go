package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

type point struct{
	z,x,y int
}

func isValid(p point, n int) bool{
	if p.z >= 0 && p.z < n && p.x >= 0 && p.x < n && p.y >= 0 && p.y < n{
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(firstLine))

	grid := make(map[point]struct{})
	var sx, sy, sz int

	for z := 0; z < n; z++ {
    	reader.ReadString('\n')
    	for x := 0; x < n; x++ {
        	line, _ := reader.ReadString('\n')
        	line = strings.TrimSpace(line)
        	for y := 0; y < len(line); y++ {
				if line[y] == '#'{
					continue
				}

            	if line[y] == 'S' {
                	sx, sy, sz = x, y, z
            	}
				grid[point{z,x,y}] = struct{}{}
        	}
    	}
	}

	start := point{sz, sx, sy}
	dist := make(map[point]int)
	dist[start] = 0
	queue := []point{start}

	dirs := [][3]int{{0,-1,0},{0,1,0},{0,0,1},{0,0,-1},{1,0,0},{-1,0,0}}

	for len(queue) > 0{
		size := len(queue)
		for i := 0; i < size; i++{
			top := queue[0]
			queue = queue[1:]
			for _, dir := range dirs{
				cz, cx, cy := top.z+dir[0], top.x+dir[1], top.y+dir[2]
				p := point{cz,cx,cy}
				if isValid(p, n){
					if _,seen := dist[p]; !seen{
						if _,ok := grid[p]; ok{
							dist[p] = dist[top]+1
							queue = append(queue, p)
						}
					}
				}
			}
		}
	}

	ans := math.MaxInt
	for k, v := range dist{
		if k.z == 0{
			ans = min(ans, v)
		}
	}

	fmt.Println(ans)
}