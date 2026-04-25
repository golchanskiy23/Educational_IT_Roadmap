/*
In this problem, a tree is an undirected graph that is connected and has no cycles.

You are given a graph that started as a tree with n nodes labeled from 1 to n, with one additional edge added. The added edge has two different vertices chosen from 1 to n, and was not an edge that already existed. The graph is represented as an array edges of length n where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the graph.

Return an edge that can be removed so that the resulting graph is a tree of n nodes. If there are multiple answers, return the answer that occurs last in the input.
*/

func find(x int, parent []int) int{
    if parent[x] != x{
        parent[x] = find(parent[x], parent)
    }

    return parent[x]
}

func union(a,b int, parent []int) bool{
    pa, pb := find(a, parent), find(b, parent)
    if pa == pb{
        return false
    }
    parent[pb] = pa
    return true
}

func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    parent := make([]int, n+1)
    for i := 1; i <= n; i++{
        parent[i] = i
    }

    for i := 0; i < n; i++{
        if !union(edges[i][0], edges[i][1], parent){
            return edges[i]
        }
    }

    return []int{}
}