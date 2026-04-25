/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.
*/

func dfs(adj [][]int, visited []int, result *[]int, curr int) bool{
    if visited[curr] == 1{return false}
    if visited[curr] == 2{return true}

    visited[curr] = 1
    for _, node := range adj[curr]{
        if !dfs(adj, visited, result, node){
            return false
        }

    }
    visited[curr] = 2
    *result = append(*result, curr)

    return true
}

func findOrder(n int, pre [][]int) []int {
    // топологическая сортировка
    // if graph[pre[i][0]] = pre[i][1]{false}
    // но не проверяет уже связные узлы
    // если не посещён - белый - 0
    // посещён, но не обработан до конца - grey - 1
    // посещён и обработан - black - 2

    adj := make([][]int, n)
    for i := 0;i < n; i++{
        adj[i] = make([]int, 0)
    }

    for i := 0; i < len(pre); i++{
        adj[pre[i][0]] = append(adj[pre[i][0]], pre[i][1])
    }

    result := make([]int, 0)
    visited := make([]int, n)
    for i := 0; i< n; i++{
        if visited[i] == 0{
            if !dfs(adj, visited, &result, i){
                return []int{}
            }
        }
    }

    return result
}